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

	scrapeTime := time.Duration(1) * time.Hour
	ticker := time.NewTicker(scrapeTime)

	fmt.Fprintln(os.Stdout, "Ecointracker api started...")

	// Runs the loop that verify the local time and reuns the functions when the specifeid time is reached
	for range ticker.C {
		t := time.Now()
		if t.Hour() == 22 {
			fmt.Fprintln(os.Stdout, "Starting a new read...")

			// Runs the GetData function of the src package, with the URL of the ecoin token deadwallet
			src.GetData(URL)

			fmt.Fprintln(os.Stdout, "End of reading step")
			fmt.Fprintln(os.Stdout, "End of cicle!")
		} else {
			fmt.Fprintln(os.Stdout, "--> Hour: ", time.Now().Format("15:04:05"))
			fmt.Fprintln(os.Stdout, "Waiting for the specified time...")
		}

	}

}
