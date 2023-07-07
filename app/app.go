package app

import (
	"github.com/urfave/cli"
)

// That function will instantiate an App from urfave/cli
func Make() *cli.App {
	app := cli.NewApp()
	app.Name = "Net Tools"
	app.Usage = "Search IP Addresses and Server names in the Internet"
	return app
}
