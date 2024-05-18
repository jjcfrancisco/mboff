package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func off(ctx *cli.Context) error {
	mbtilesFilepath := ctx.Args().Get(0)
	if !fileExists(mbtilesFilepath) {
		log.Fatal(fmt.Sprintf("MBTiles file path in %s does not exist.", mbtilesFilepath))
	}
	conn, err := createConn(mbtilesFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.db.Close()
	rawKeyValue := ctx.Args().Get(1)
	keyValue, err := parseKeyValue(rawKeyValue)
	if err != nil {
		log.Fatal(err)
	}
	results, err := conn.find(keyValue)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) != 0 {
		if err := conn.update(results); err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Successful. The key/value %s was found in %d instance(s). MBTiles updated.", rawKeyValue, len(results)))
	} else {
		log.Fatal(fmt.Sprintf("The key value %s was not present in the MBTiles file. Nothing was removed.", keyValue))
	}

	return nil

}

func startCli() *cli.App {

	// Version Flag
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "Print out version",
	}

	// Main CLI
	app := &cli.App{
		Name: "Mboff",
		// Version: "v0.0.1",
		Usage:  "Mboff optimises your mbtiles by removing unnecessary data.",
		Action: off,
	}

	return app

}
