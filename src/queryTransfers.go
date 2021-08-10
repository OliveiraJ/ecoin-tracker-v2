package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type transfers struct {
	Count    int64
	Days     int
	Min_date string
	Max_date string
}

type bitqueryData struct {
	Data struct {
		Ethereum struct {
			Transfers []transfers
		}
	}
}

var b bitqueryData

func GetTransfers() transfers {

	jsonData := map[string]string{
		"query": `
			query ($network: EthereumNetwork!, $token:String!, $from: ISO8601DateTime, $till: ISO8601DateTime) {
				ethereum(network: $network) {
					transfers(currency: {is: $token}, amount: {gt: 0}, date: {since: $from, till: $till}) {
						count
						days: count(uniq: dates)
						min_date: minimum(of: date)
						max_date: maximum(of: date)
					}
				}
			}
		`,
		"variables": `
			{
				"limit":10,
				"offset":0,
				"network":"bsc",
				"token":"0x4cbdfad03b968bf43449d0908f319ae4a5a33371"
			}
		`,
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://graphql.bitquery.io", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	req.Header.Add("content-type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &b)

	return b.Data.Ethereum.Transfers[0]

}
