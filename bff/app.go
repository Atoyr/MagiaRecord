package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/atoyr/MagiaRecord/bff/handler"
	"github.com/labstack/echo"
	"github.com/urfave/cli/v2"
)

var server string
var instance string
var user string
var password string
var db string
var port string

var okstring = "[\x1b[32m OK \x1b[0m]"
var failstring = "[\x1b[31mFAIL\x1b[0m]"
var infostring = "[\x1b[36mINFO\x1b[0m]"

func main() {
	serverFlag := &cli.StringFlag{
		Name:        "server",
		Aliases:     []string{"s"},
		Value:       "",
		Usage:       "SQLServer Server Name",
		EnvVars:     []string{"DBSERVER"},
		Destination: &server,
		Required:    true,
	}
	instanceFlag := &cli.StringFlag{
		Name:        "instance",
		Aliases:     []string{"i"},
		Value:       "",
		Usage:       "SQLServer Server Instance Name",
		EnvVars:     []string{"DBINSTANCE"},
		Destination: &instance,
	}
	userFlag := &cli.StringFlag{
		Name:        "user",
		Aliases:     []string{"u"},
		Value:       "sa",
		Usage:       "SQLServer Server User",
		EnvVars:     []string{"DBUSER"},
		Destination: &user,
	}
	passwordFlag := &cli.StringFlag{
		Name:        "password",
		Aliases:     []string{"p"},
		Value:       "",
		Usage:       "SQLServer Server Password",
		EnvVars:     []string{"DBPASS"},
		Destination: &password,
		Required:    true,
	}
	databaseFlag := &cli.StringFlag{
		Name:        "database",
		Aliases:     []string{"d"},
		Value:       "master",
		Usage:       "SQLServer Server using database",
		Destination: &db,
	}
	portFlag := &cli.StringFlag{
		Name:        "port",
		Value:       ":80",
		Usage:       "Web Api Port",
		Destination: &port,
	}

	app := new(cli.App)
	app.Name = "Magia Record Backend For Frontend"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		serverFlag,
		instanceFlag,
		userFlag,
		passwordFlag,
		databaseFlag,
		portFlag,
	}

	app.Action = func(context *cli.Context) error {
		var err error
		fmt.Printf("%s Start Check Database Is Executing.\n", infostring)
		err = checkExecuteDatabase()
		if err != nil {
			return fmt.Errorf("SQL Server Connect Error : %s", err)
		}
		fmt.Printf("%s SQL Server Connect\n", okstring)
		e := echo.New()

		e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				con := handler.Context{Context: c, DriverName: "", DataSourceName: ""}
				return h(&con)
			}
		})
		e.GET("/magical_girls", c(handler.GetMagicalGirlsHandler()))
		e.GET("/types", c(handler.GetTypesHandler()))
		e.GET("/attributes", c(handler.GetAttributesHandler()))

		fmt.Printf("%s Start Echo Web API\n", infostring)
		e.Start(port)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s %s\n", failstring, err.Error())
	}
}

func c(h handler.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*handler.Context))
	}
}

func getConnectionString() string {
	var ret = make([]byte, 0, 512)
	ret = append(ret, "server="...)
	ret = append(ret, server...)
	if instance != "" {
		ret = append(ret, "\\"...)
		ret = append(ret, instance...)
	}
	ret = append(ret, ";user id="...)
	ret = append(ret, user...)
	ret = append(ret, ";password="...)
	ret = append(ret, password...)
	ret = append(ret, ";database="...)
	ret = append(ret, db...)
	return string(ret)
}

func checkExecuteDatabase() error {
	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
