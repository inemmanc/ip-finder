package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Generate will return CLI app ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "IP FINDER APPLICATION"
	app.Usage = "Find the IP and NAME of a server"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			// default value
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find IP adresses on the internet",
			Flags:  flags,
			Action: ipFinder,
		},
		{
			Name:   "server",
			Usage:  "Find server names on the internet",
			Flags:  flags,
			Action: nameFinder,
		},
	}

	return app
}

func ipFinder(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		fmt.Println("ERROR FINDING IP")
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func nameFinder(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		fmt.Println("ERROR NAME SERVER")
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
