package main

import (
	"fmt"
	"ipcli/app"
	"os"
)

func main() {
	fmt.Println("START\n")
	appGNR := app.Generate()
	err := appGNR.Run(os.Args)

	if err != nil {
		fmt.Println("ERROR")
	}
}
