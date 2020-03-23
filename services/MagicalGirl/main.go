package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "MagiaRecord MagicalGirl Service"
	app.Usage = "MagiaRecord MagicalGirl Service"
	app.Version = version

	var httpPort uint

	app.Flags = []cli.Flag{
		&cli.UintFlag{
			Name:        "HttpPort",
			Aliases:     []string{"p"},
			Value:       8080,
			EnvVars:     []string{"HTTPPORT"},
			Destination: &httpPort,
			Required:    false,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
