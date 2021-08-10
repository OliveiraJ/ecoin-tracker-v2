package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Stores the data received by the query.
type transfers struct {
	Count    int64
	Days     int
	Min_date string
	Max_date string
}

// Works as a adapter so the received data, once this has a particular structure that demands the data to be wraped by other structs
// dificulting it conversion.
type bitqueryData struct {
	Data struct {
		Ethereum struct {
			Transfers []transfers
		}
	}
}

var b bitqueryData

// Do a query passing the schema of a graphql query and its variables to the https://graphql.bitquery.io api, returning a int64 value with the ammount of transfers
// made with the Ecoin Finance token.
func GetTransfers() transfers {

	// Map that stores the Schema of the query, in a string format.
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

	// Converts the JsonData map to the Json format, returning a []Bytes and a Error.
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// Requests data to the https://graphql.bitquery.io , a Graphql API that returs data about Cryptocoins passing the jsonValue
	// variable as the body of the request, through a POST method.
	req, err := http.NewRequest("POST", "https://graphql.bitquery.io", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	// Add parammeters to the Header of the request.
	req.Header.Add("content-type", "application/json")

	// Creates a new Http Client.
	client := &http.Client{}

	// Do the http request, returning a http response and an error
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Closes the body of the response at the end of the function
	defer response.Body.Close()

	// Read the body of the response, returning a []bytes and an error
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// Converts the []bytes in the json format to the bitqueryData struct format, returning an error.
	json.Unmarshal(data, &b)

	return b.Data.Ethereum.Transfers[0]

}
