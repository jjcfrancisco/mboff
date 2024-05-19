package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/urfave/cli/v2"
)

func off(ctx *cli.Context) error {
	// MBTiles file path
	mbtilesFilepath := ctx.Args().Get(0)
	if !fileExists(mbtilesFilepath) {
		log.Fatal(fmt.Sprintf("MBTiles file path in %s does not exist.", mbtilesFilepath))
	}
	conn, err := createConn(mbtilesFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.db.Close()
	// Key Value
	rawKeyValue := ctx.Args().Get(1)
	keyValue, err := parseKeyValue(rawKeyValue)
	if err != nil {
		log.Fatal(err)
	}
	// Zoom Level
	strZoomLevel := ctx.Args().Get(2)
	var intZoomLevel *int
	if strZoomLevel == "" {
		intZoomLevel = nil
	} else {
		num, err := strconv.Atoi(strZoomLevel)
		if err != nil {
			log.Fatal(err)
		}
		intZoomLevel = &num
	}
	// Find feature(s)
	results, err := conn.find(keyValue, intZoomLevel)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) != 0 {
		// Update feature(s)
		if err := conn.update(results, intZoomLevel); err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Successful. The key/value %s was found. MBTiles updated.", rawKeyValue))
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
		Version: "v0.2.2",
		Usage:  "Mboff optimises your mbtiles by removing unnecessary data.",
		Action: off,
	}

	return app

}

