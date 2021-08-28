package src

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

var holders int64

// Scrapes and return the numeber of holders of the ecoin token, from the Bscscan.com and retunrs it as a int64
func GetHolders(URL string) int64 {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	CollectorHolders := colly.NewCollector()

	// CollectorHolders scrapes the Bscscan.com the ammount holders of the Ecoin Finance token
	CollectorHolders.OnHTML("#ContentPlaceHolder1_tr_tokenHolders", func(e *colly.HTMLElement) {
		r := re.FindString(e.Text)

		//Replace the comma carachter from the string after the regex treatment making it easy to convert to a number later.
		r = strings.Replace(r, comma, replaceArgument, replaceTimes)

		// Converts the screaped value to a string and adds it to the Read.Holders property
		if r != "" {
			value, err := strconv.ParseInt(r, 10, 64)
			if err == nil {
				holders = value
			}
		}
	})

	// Gives the collectors a start and passes the URL it should visit
	CollectorHolders.Visit(URL)

	return holders
}
