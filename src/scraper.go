package src

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const comma string = ","
const replaceArgument string = ""
const replaceTimes int = 4

type Read struct {
	BurnedTokens float64 `json:"burnedTokens"`
	Holders      int64   `json:"holders"`
	Transfers    int64   `json:"transfers"`
	Date         string  `json:"date"`
}

var AllReads []Read

func GetData(URL string) {

	readJson()

	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	CollectorBalance := colly.NewCollector()

	CollectorHolders := colly.NewCollector()

	var Read Read

	CollectorBalance.OnHTML("#ContentPlaceHolder1_divFilteredHolderBalance", func(e *colly.HTMLElement) {

		r := re.FindString(e.Text)
		//Replace the comma carachter from the string after the regex treatment making it easy to convert to a number later.
		r = strings.Replace(r, comma, replaceArgument, replaceTimes)

		//Adding the time value to the Read.Date property.
		Read.Date = time.Now().Format("2006/01/02")

		if r != "" {
			value, err := strconv.ParseFloat(r, 10)
			if err == nil {
				Read.BurnedTokens = value
			}
		}
	})

	//cHolders capture the number of holders of e-coin
	CollectorHolders.OnHTML("#ContentPlaceHolder1_tr_tokenHolders", func(e *colly.HTMLElement) {
		r := re.FindString(e.Text)

		//Replace the comma carachter from the string after the regex treatment making it easy to convert to a number later.
		r = strings.Replace(r, comma, replaceArgument, replaceTimes)

		if r != "" {
			value, err := strconv.ParseInt(r, 10, 64)
			if err == nil {
				Read.Holders = value
			}
		}
	})
	CollectorBalance.OnRequest(func(request *colly.Request) {
		log.Println("Visiting", request.URL.String())
	})

	CollectorBalance.Visit(URL)
	CollectorHolders.Visit(URL)

	AllReads = append(AllReads, Read)

	writeJSON(AllReads)
	convertJSON()
}
func readJson() []Read {
	if !Exists(`./data/data.json`) {
		log.Println("Criando arquivo JSON")
		jsonFile, err := os.Create(`./data/data.json`)
		if err != nil {
			panic(err)
		}
		log.Println("Arquivo JSON criado")
		defer jsonFile.Close()
	}
	log.Println("Lendo JSON")

	jsonFile, err := os.Open(`./data/data.json`)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValueJSON, &AllReads)
	return AllReads
}
func writeJSON(data []Read) {
	log.Println("Salvando dados")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	err = ioutil.WriteFile("./data/data.json", file, 0644)
	if err != nil {
		log.Println(err)
	}

}
func convertJSON() {
	readJson()
	log.Println("Escrevendo arquivo CSV")
	csvFile, err := os.Create("./data/data.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	headRow := []string{"Burned Tokens", "Holders", "Date"}
	writer.Write(headRow)
	for _, jsonData := range AllReads {
		var row []string
		row = append(row, fmt.Sprintf("%f", jsonData.BurnedTokens))
		row = append(row, fmt.Sprint(jsonData.Holders))
		row = append(row, jsonData.Date)
		writer.Write(row)
	}

	writer.Flush()
}
func Exists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
