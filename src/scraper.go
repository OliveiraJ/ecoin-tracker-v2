package src

import (
	"github.com/OliveiraJ/ecoin-tracker-v2/db"
	"github.com/OliveiraJ/ecoin-tracker-v2/models"
	src "github.com/OliveiraJ/ecoin-tracker-v2/src/collectors"
	utils "github.com/OliveiraJ/ecoin-tracker-v2/src/utils"
)

const pathFileJson = "./data/data.json"
const pathFileCsv = "./data/data.csv"
const pathFolder = "./data"

// Slice of the type Read returned by the GetData function
var AllReads []models.Read

// Calls the GetBalance, GetHolders and GetTransfers saving the resulting data in a .JSON and a .CSV files
// wrapping all the process in acquiring all the data and populate the Json file, CSV file and PostgreSQL database
func GetData(URL string) {
	// Creates a connection with the database
	database := db.NewDbConnection()
	// Creates a connection with the test_database
	test_database := db.NewTestDbConnection()

	AllReads = utils.ReadJson(pathFolder, pathFileJson, AllReads)

	// Auxiliar variable Read of type Read.
	var Read models.Read

	Read.BurnedTokens, Read.Date, Read.Hour = src.GetBalance(URL)
	Read.Holders = src.GetHolders(URL)
	Read.Transfers = src.GetTransfers().Count

	// Increments the ID value to match the id of the last entry on the Data.json file, sincronizing the file with the database
	if len(AllReads) == 0 {
		Read.ID = 1
	} else {
		Read.ID = AllReads[len(AllReads)-1].ID + 1
	}

	// Updates the AllReads slice with the most recent Read
	AllReads = append(AllReads, Read)

	utils.WriteJSON(AllReads, pathFileJson)

	// If database doesn't habe a models.Read table, then creates a table and inserts in it all the values in the Json File
	// otherwise inserts only the actual Read.
	if !database.Debug().Migrator().HasTable(&models.Read{}) {
		db.Setup(database)
		database.Debug().CreateInBatches(AllReads, len(AllReads))
	} else {
		database.Debug().Create(&Read)
	}

	// Repeat the process above to the test_database
	if !test_database.Debug().Migrator().HasTable(&models.Read{}) {
		db.Setup(test_database)
		test_database.Debug().CreateInBatches(AllReads, len(AllReads))
	} else {
		test_database.Debug().Create(&Read)
	}

	utils.ConvertJSON(pathFileCsv, pathFolder, pathFileJson, AllReads)

}
