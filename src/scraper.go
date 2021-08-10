package src

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	Hour         string  `json:"hour"`
}

// Slice of the type Read returned by the GetData function
var AllReads []Read

// GetData scrapes from Bscscan.com, using the Colly package, a Regex and the Replace function, readig the Data.json file
// and writing it with the new data, preserving the previous data in it. Also calls the ConvertJson function, generating
// a Data.csv by converting the Data.json file to it.
func GetData(URL string) {

	// Calls the ReadJson function, reading the Data.json file so it can be updated with the new scraped data.
	ReadJson()

	// Regular expression that helps to clean the scraped data.
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	// Collector from colly package.
	CollectorBalance := colly.NewCollector()
	CollectorHolders := colly.NewCollector()

	// Auxiliar variable Read of type Read.
	var Read Read

	// CollectorBalance scrapes the Bscscan.com the ammount of burned tokens in the dead wallet and adds the time.
	CollectorBalance.OnHTML("#ContentPlaceHolder1_divFilteredHolderBalance", func(e *colly.HTMLElement) {

		// Aplys the regex Re to the e.Text scraped by the collector, helpíng to clean it and keep only the usefull data
		r := re.FindString(e.Text)

		// Replace the comma carachter from the string after the regex treatment making it easy to convert to a number later.
		r = strings.Replace(r, comma, replaceArgument, replaceTimes)

		// Adds the time value to the Read.Date property.
		Read.Date = time.Now().Format("02/01/2006")

		// Adds the time value to the Read.Hour property
		Read.Hour = time.Now().Format("15:04:05")

		// Converts the screaped value to a string and adds it to the Read.BurnedTokens property
		if r != "" {
			value, err := strconv.ParseFloat(r, 10)
			if err == nil {
				Read.BurnedTokens = value
			}
		}
	})

	// CollectorHolders scrapes the Bscscan.com the ammount holders of the Ecoin Finance token
	CollectorHolders.OnHTML("#ContentPlaceHolder1_tr_tokenHolders", func(e *colly.HTMLElement) {
		r := re.FindString(e.Text)

		//Replace the comma carachter from the string after the regex treatment making it easy to convert to a number later.
		r = strings.Replace(r, comma, replaceArgument, replaceTimes)

		// Converts the screaped value to a string and adds it to the Read.Holders property
		if r != "" {
			value, err := strconv.ParseInt(r, 10, 64)
			if err == nil {
				Read.Holders = value
			}
		}
	})

	CollectorBalance.OnRequest(func(request *colly.Request) {
		fmt.Fprintln(os.Stdout, "Visiting", request.URL.String())
	})

	// Gives the collectors a start and passes the URL it should visit
	CollectorBalance.Visit(URL)
	CollectorHolders.Visit(URL)

	Read.Transfers = GetTransfers().Count

	AllReads = append(AllReads, Read)

	writeJSON(AllReads)
	convertJSON()

}

// ReadJson reads the JSON file and returns a slice of type Read
func ReadJson() []Read {

	//Verify if the data.json file exists and creat a new one if it doesnt
	if !Exists(`/home/jordan/Documentos/EcoinTracker/data.json`) {
		fmt.Fprintln(os.Stdout, "Criando arquivo JSON")
		jsonFile, err := os.Create(`/home/jordan/Documentos/EcoinTracker/data.json`)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stdout, "Arquivo JSON criado")
		defer jsonFile.Close()
	}
	fmt.Fprintln(os.Stdout, "Lendo JSON")

	//Open data.json file
	jsonFile, err := os.Open(`/home/jordan/Documentos/EcoinTracker/data.json`)
	if err != nil {
		panic(err)
	}

	// closes the data.json file
	defer jsonFile.Close()

	byteValueJSON, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Erro ao ler arquivo JSON --> ", err)
	}

	json.Unmarshal(byteValueJSON, &AllReads)

	// Return the slice with the read data from the scraper and json file
	return AllReads
}

// WriteJson write the data in a JSON file
func writeJSON(data []Read) {
	fmt.Fprintln(os.Stdout, "Salvando dados")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stdout, "Unable to create json file")
		return
	}

	err = ioutil.WriteFile("/home/jordan/Documentos/EcoinTracker/data.json", file, 0644)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

}

// CovertJson converts  the Data.json file in the Data.csv file
func convertJSON() {
	ReadJson()
	fmt.Println("Escrevendo arquivo CSV")
	csvFile, err := os.Create("/home/jordan/Documentos/EcoinTracker/data.csv")
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
