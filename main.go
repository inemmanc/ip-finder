package main

import (
	"fmt"
	"ipcli/app"
	"log"
	"os"
)

func main() {
	fmt.Println("-_-_-_-_ IP and SERVER FINDER _-_-_-_-")
	fmt.Println()
	appGNR := app.Generate()
	err := appGNR.Run(os.Args)

	if err != nil {
		fmt.Println("==== ERROR ====")
		log.Fatal(err)
	}
}
