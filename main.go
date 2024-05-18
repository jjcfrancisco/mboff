package main

import (
	"log"
	"os"
)

func main() {
	app := startCli()
	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
