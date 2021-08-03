package main

import (
	"fmt"
	"os"
	"time"

	"github.com/OliveiraJ/ecoin-tracker-v2/server"
	"github.com/OliveiraJ/ecoin-tracker-v2/src"
)

const URL string = "https://bscscan.com/token/0x4cbdfad03b968bf43449d0908f319ae4a5a33371?a=0x000000000000000000000000000000000000dead"

func main() {
	// go server.GetJson()
	go server.HandleResquests()

	scrapeTime := time.Duration(1) * time.Minute
	ticker := time.NewTicker(scrapeTime)

	//ticker, reprete a função a cada intervalo de tempo.
	fmt.Fprintln(os.Stdout, "EcoinTracker iniciado...")
	for range ticker.C {
		t := time.Now()
		if t.Hour() == 12 {
			fmt.Fprintln(os.Stdout, "Começando nova leitura...")
			src.GetData(URL)
			fmt.Fprintln(os.Stdout, "Leitura efetuda")
			fmt.Fprintln(os.Stdout, "Fim do Ciclo!")
		} else {
			fmt.Fprintln(os.Stdout, "Aguardando horário especificado...")
			fmt.Fprintln(os.Stdout, "--> Hora atual: ", time.Now().Hour())
		}

	}

}
