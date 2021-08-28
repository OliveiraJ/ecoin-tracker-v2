package src

import (
	"fmt"
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

var date, hour string
var burnedTokens float64

// Scrapes the balance of the dead wallet on Bscscan.com and return its value as a float64, also returnin the date and hour
// of the read in the dd/mm/yyyy and hh:mm:ss format
func GetBalance(URL string) (float64, string, string) {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	CollectorBalance := colly.NewCollector()

	// CollectorBalance scrapes the Bscscan.com the ammount of burned tokens in the dead wallet and adds the time.
	CollectorBalance.OnHTML("#ContentPlaceHolder1_divFilteredHolderBalance", func(e *colly.HTMLElement) {

		// Aplys the regex Re to the e.Text scraped by the collector, helpíng to clean it and keep only the usefull data
		r := re.FindString(e.Text)

		// Replace the comma carachter from the string after the regex treatment making it easy to convert to a number later.
		r = strings.Replace(r, comma, replaceArgument, replaceTimes)

		// Adds the time value to the Read.Date property.
		date = time.Now().Format("02/01/2006")

		// Adds the time value to the Read.Hour property
		hour = time.Now().Format("15:04:05")

		// Converts the screaped value to a string and adds it to the Read.BurnedTokens property
		if r != "" {
			value, err := strconv.ParseFloat(r, 10)
			if err == nil {
				burnedTokens = value
			}
		}
	})

	CollectorBalance.OnRequest(func(request *colly.Request) {
		fmt.Fprintln(os.Stdout, "Visiting", request.URL.String())
	})

	// Gives the collectors a start and passes the URL it should visit
	CollectorBalance.Visit(URL)

	return burnedTokens, date, hour
}
