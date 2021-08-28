package src

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"os"

	"github.com/OliveiraJ/ecoin-tracker-v2/models"
	src "github.com/OliveiraJ/ecoin-tracker-v2/src/collectors"
)

const comma string = ","
const replaceArgument string = ""
const replaceTimes int = 4
const pathFileJson = "./data/data.json"
const pathFileCsv = "./data/data.csv"
const pathFolder = "./data"

// Slice of the type Read returned by the GetData function
var AllReads []models.Read

// Calls the GetBalance, GetHolders and GetTransfers saving the resulting data in a .JSON and a .CSV files
func GetData(URL string) {

	// Calls the ReadJson function, reading the Data.json file so it can be updated with the new scraped data.
	ReadJson()

	// Auxiliar variable Read of type Read.
	var Read models.Read

	Read.BurnedTokens, Read.Date, Read.Hour = src.GetBalance(URL)
	Read.Holders = src.GetHolders(URL)
	// Calls the GetTransfers function, storing the returned value in the Read.Transfers property
	Read.Transfers = src.GetTransfers().Count

	// Updates the AllReads slice with the most recent Read
	AllReads = append(AllReads, Read)

	// Writes the data.json file with the most updates AllReads slice as its content.
	writeJSON(AllReads)
	convertJSON()
}

// ReadJson reads the JSON file and returns a slice of type Read
func ReadJson() []models.Read {

	// Verify if the data diretory exists and creat it if it doesnt
	if !Exists(pathFolder) {
		fmt.Fprintln(os.Stdout, "./data folder doesn't exist, creating folder")
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Fprintln(os.Stdout, "Error on creating ./data folder")
			panic(err)
		}
		fmt.Fprintln(os.Stdout, "./data folder created")
	}

	// Verify if the data.json file exists and creat a new one if it doesnt
	if !Exists(pathFileJson) {
		fmt.Fprintln(os.Stdout, "data.json doesn't exist, creating data.json file")
		jsonFile, err := os.Create(pathFileJson)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stdout, "data.json file created")
		defer jsonFile.Close()
	}
	fmt.Fprintln(os.Stdout, "Reading data.json file")

	//Open data.json file
	jsonFile, err := os.Open(pathFileJson)
	if err != nil {
		panic(err)
	}

	// closes the data.json file
	defer jsonFile.Close()

	byteValueJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Error on reading json file --> ", err)
	}

	json.Unmarshal(byteValueJSON, &AllReads)

	// Return the slice with the read data from the scraper and json file
	return AllReads
}

// WriteJson write the data in a JSON file
func writeJSON(data []models.Read) {
	fmt.Fprintln(os.Stdout, "Saving data")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stdout, "Unable to create json file")
		return
	}

	err = ioutil.WriteFile(pathFileJson, file, 0644)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

}

// CovertJson converts  the Data.json file in the Data.csv file
func convertJSON() {
	ReadJson()
	fmt.Println("Writing data.csv file")
	csvFile, err := os.Create(pathFileCsv)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	headRow := []string{"Burned Tokens", "Holders", "Transfers", "Date", "Hour"}
	writer.Write(headRow)
	for _, jsonData := range AllReads {
		var row []string
		row = append(row, fmt.Sprintf("%f", jsonData.BurnedTokens))
		row = append(row, fmt.Sprint(jsonData.Holders))
		row = append(row, fmt.Sprint(jsonData.Transfers))
		row = append(row, jsonData.Date)
		row = append(row, jsonData.Hour)
		writer.Write(row)
	}

	writer.Flush()
}

// Exists verifys if the json file exists and returns false if it doesn't or true if it does
func Exists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
