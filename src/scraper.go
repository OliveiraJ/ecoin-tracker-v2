package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

	WriteJSON(AllReads)

	log.Println(Read)
}
func WriteJSON(data []Read) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}

	_ = ioutil.WriteFile("data.json", file, 0644)
}
