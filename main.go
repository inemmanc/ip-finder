package main

import (
	"fmt"
	"ipcli/app"
	"os"
)

func main() {
	fmt.Println("START")
	appGNR := app.Generate()
	err := appGNR.Run(os.Args)
	if err != nil {
		fmt.Println("ERROR")
	}
}
