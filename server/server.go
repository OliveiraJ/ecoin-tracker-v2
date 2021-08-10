package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/OliveiraJ/ecoin-tracker-v2/src"
	"github.com/gorilla/mux"
)

const Dir = "/home/jordan/Documentos/EcoinTracker"

// GetJson function returns the JSON file in the respective route
func GetJson(res http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&res, req)
	http.ServeFile(res, req, Dir+"/data.json")
}

// GetCSV function returns the CSV file the the respective route
func GetCsv(res http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&res, req)
	http.ServeFile(res, req, Dir+"/data.csv")
}

// Get function the data in a JSON format
func Get(res http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&res, req)
	log.Println("Retornando JSON")
	json.NewEncoder(res).Encode(src.ReadJson())
}

// HandleRequest function handles the routes of the application using the github.com/gorilla/mux package
func HandleResquests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Get)
	router.HandleFunc("/csv", GetCsv)
	router.HandleFunc("/json", GetJson)

	log.Fatal(http.ListenAndServe(":9000", router))
}

func setupCorsResponse(res *http.ResponseWriter, req *http.Request) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(*res).Header().Set("Access-Control-Allow-Headres", "Accept, Content-Type, Content-Length, Autohrization")
	(*res).Header().Set("Content-Type", "application/json")
}
