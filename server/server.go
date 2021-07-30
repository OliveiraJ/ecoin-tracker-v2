package server

import (
	"log"
	"net/http"
)

func GetJson() {
	fs := http.FileServer(http.Dir("/home/jordan/Documentos/EcoinTracker"))
	log.Fatal(http.ListenAndServe(":9000", fs))

}
