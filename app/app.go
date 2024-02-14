package app

import (
	"github.com/urfave/cli"
)

// Generate will return CLI app ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP FINDER APPLICATION"
	app.Usage = "Find the IP and NAME of a server"

	return app
}
