package main

import (
	"log"
	//"net/http"
	"path/filepath"
	"time"

	"github.com/OliveiraJ/ecoin-tracker-v2/src"
)

const URL string = "https://bscscan.com/token/0x4cbdfad03b968bf43449d0908f319ae4a5a33371?a=0x000000000000000000000000000000000000dead"

var DirData = filepath.FromSlash("data")

func main() {
	scrapeTime := time.Duration(1) * time.Hour
	ticker := time.NewTicker(scrapeTime)

	// fs := http.FileServer(http.Dir("/home/jordan/Documentos/EcoinTracker"))
	// log.Fatal(http.ListenAndServe("192.168.0.125:9000", fs))

	// http.HandleFunc("/csv", func(res http.ResponseWriter, req *http.Request) {
	// 	http.ServeFile(res, req, filepath.Join(DirData, "/data.csv"))
	// })
	// log.Fatal(http.ListenAndServe("192.168.0.125:9000", nil))

	//ticker, reprete a função a cada intervalo de tempo.
	for range ticker.C {
		t := time.Now()
		if t.Hour() == 22 {
			log.Println("Começando nova leitura...")
			src.GetData(URL)
			log.Println("Leitura efetuda")
			log.Println("Fim do Ciclo!")
		} else {
			log.Println("Aguardando horário especificado...")
		}

	}

}
