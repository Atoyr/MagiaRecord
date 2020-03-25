package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mg "github.com/atoyr/MagiaRecord/services/Attribute/app"
	_ "github.com/atoyr/MagiaRecord/services/Attribute/app/controllers"
	"github.com/urfave/cli/v2"
)

const version = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "MagiaRecord Attribute Service"
	app.Usage = "MagiaRecord Attribute Service"
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

	app.Action = func(c *cli.Context) error {
		conf := mg.NewConfig()
		conf.HttpPort = httpPort
		conf.DatabaseServer = databaseServer
		conf.DatabasePort = databasePort
		conf.DatabaseUser = databaseUser
		conf.DatabasePassword = databasePassword
		conf.UsingDatabase = usingDatabase

		e := http.ListenAndServe(fmt.Sprintf(":%d", conf.HttpPort), nil)
		if e != nil {
			log.Println(e)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
