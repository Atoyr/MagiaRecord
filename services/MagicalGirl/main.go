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
	var databaseServer string
	var databasePort uint
	var databaseUser string
	var databasePassword string
	var usingDatabase string

	app.Flags = []cli.Flag{
		&cli.UintFlag{
			Name:        "http-port",
			Aliases:     []string{"H"},
			Value:       8080,
			EnvVars:     []string{"HTTP_PORT"},
			Destination: &httpPort,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "db-server",
			Aliases:     []string{"s"},
			Value:       "",
			EnvVars:     []string{"DB_SERVER"},
			Destination: &databaseServer,
			Required:    true,
		},
		&cli.UintFlag{
			Name:        "db-port",
			Aliases:     []string{"P"},
			Value:       3306,
			EnvVars:     []string{"DB_PORT"},
			Destination: &databasePort,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "db-user",
			Aliases:     []string{"u"},
			Value:       "root",
			EnvVars:     []string{"DB_USER"},
			Destination: &databaseUser,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "db-password",
			Aliases:     []string{"p"},
			Value:       "",
			EnvVars:     []string{"DB_PASSWORD"},
			Destination: &databasePassword,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "using-db",
			Aliases:     []string{"U"},
			Value:       "",
			EnvVars:     []string{"USING_DATABASE"},
			Destination: &usingDatabase,
			Required:    true,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
