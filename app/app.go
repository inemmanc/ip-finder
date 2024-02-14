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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find IP adresses on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					// default value
					Value: "google.com",
				},
			},
			Action: ipFinder,
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
