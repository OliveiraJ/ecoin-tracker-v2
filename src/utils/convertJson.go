package src

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/OliveiraJ/ecoin-tracker-v2/models"
)

// CovertJson converts  the Data.json file in the Data.csv file
func ConvertJSON(pathFileCsv string, pathFolder string, pathFileJson string, AllReads []models.Read) {
	ReadJson(pathFolder, pathFileJson, AllReads)
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
