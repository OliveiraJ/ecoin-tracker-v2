package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/OliveiraJ/ecoin-tracker-v2/models"
)

// WriteJson write the data in a JSON file
func WriteJSON(data []models.Read, pathFileJson string) {
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
