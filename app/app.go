package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// That function will instantiate an App from urfave/cli
func Make() *cli.App {
	app := cli.NewApp()
	app.Name = "Net Tools"
	app.Usage = "Search IP Addresses and Server names in the Internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find the IP Address of given a host",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "github.com",
				},
			},
			Action: findIps,
		},
	}
	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
