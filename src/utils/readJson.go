package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/OliveiraJ/ecoin-tracker-v2/models"
)

// ReadJson reads the JSON file and returns a slice of type Read
func ReadJson(pathFolder string, pathFileJson string, AllReads []models.Read) []models.Read {

	// Verify if the data diretory exists and creat it if it doesnt
	if !Exists(pathFolder) {
		fmt.Fprintln(os.Stdout, "./data folder doesn't exist, creating folder")
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Fprintln(os.Stdout, "Error on creating ./data folder ", err)
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

// Exists verifys if the json file exists and returns false if it doesn't or true if it does
func Exists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
