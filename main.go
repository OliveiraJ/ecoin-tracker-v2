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
	// Goroutine that runs the server functions
	go server.HandleResquests()

	scrapeTime := time.Duration(5) * time.Second
	ticker := time.NewTicker(scrapeTime)

	fmt.Fprintln(os.Stdout, "EcoinTracker iniciado...")

	// Runs the loop that verify the local time and reuns the functions when the specifeid time is reached
	for range ticker.C {
		t := time.Now()
		if t.Hour() == 22 {
			fmt.Fprintln(os.Stdout, "Começando nova leitura...")

			// Runs the GetData function of the src package, with the URL of the ecoin token deadwallet
			src.GetData(URL)

			fmt.Fprintln(os.Stdout, "Leitura efetuda")
			fmt.Fprintln(os.Stdout, "Fim do Ciclo!")
		} else {
			fmt.Fprintln(os.Stdout, "--> Hora atual: ", time.Now().Format("15:04:05"))
			fmt.Fprintln(os.Stdout, "Aguardando horário especificado...")
		}

	}

}
