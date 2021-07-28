package main

import (
	"log"
	"time"

	"github.com/OliveiraJ/ecoin-tracker-v2/src"
)

const URL string = "https://bscscan.com/token/0x4cbdfad03b968bf43449d0908f319ae4a5a33371?a=0x000000000000000000000000000000000000dead"

func main() {
	scrapeTime := time.Duration(1) * time.Minute
	ticker := time.NewTicker(scrapeTime)

	for range ticker.C {
		log.Println("Começando nova leitura...")
		src.GetData(URL)
		log.Println("Leitura efetuda")

	}
}
