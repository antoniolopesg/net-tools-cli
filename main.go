package main

import (
	"log"
	"os"

	"github.com/antoniolopesg/net-tools-cli/app"
)

func main() {
	application := app.Make()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
