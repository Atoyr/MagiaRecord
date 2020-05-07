package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/urfave/cli/v2"
)

var server string
var instance string
var user string
var password string
var db string
var port string
var folder string

var okstring = "[\x1b[32m OK \x1b[0m]"
var failstring = "[\x1b[31mFAIL\x1b[0m]"
var infostring = "[\x1b[34mINFO\x1b[0m]"

type version struct {
	mejor    int
	miner    int
	revision int
}

func (v1 *version) compare(v2 version) int {
	if v1.mejor < v2.mejor {
		return -1
	} else if v1.mejor > v2.mejor {
		return 1
	}

	if v1.miner < v2.miner {
		return -1
	} else if v1.miner > v2.miner {
		return 1
	}

	if v1.revision < v2.revision {
		return -1
	} else if v1.revision > v2.revision {
		return 1
	}

	return 0
}

func (v *version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.mejor, v.miner, v.revision)
}

func main() {
	current, _ := os.Getwd()
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
	folderFlag := &cli.StringFlag{
		Name:        "folder",
		Aliases:     []string{"f"},
		Value:       current,
		Usage:       "SQL Query Current Folder",
		Destination: &folder,
	}

	app := new(cli.App)
	app.Name = "Database Version Manager"
	app.Commands = []*cli.Command{
		{
			Name:    "CheckDatabase",
			Aliases: []string{"c"},
			Usage:   "Check Database Version",
			Flags: []cli.Flag{
				serverFlag,
				instanceFlag,
				userFlag,
				passwordFlag,
				databaseFlag,
			},
			Action: checkDatabaseAction,
		},
		{
			Name:    "UpdateDatabase",
			Aliases: []string{"u"},
			Usage:   "Update Database to Version",
			Flags: []cli.Flag{
				serverFlag,
				instanceFlag,
				userFlag,
				passwordFlag,
				databaseFlag,
				folderFlag,
			},
			Action: updateDatabaseAction,
		},
		{
			Name:    "QueryVersion",
			Aliases: []string{"qv"},
			Usage:   "Get Query Version",
			Flags: []cli.Flag{
				folderFlag,
			},
			Action: queryVersionAction,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%s %s\n", failstring, err.Error())
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

func checkDatabaseAction(c *cli.Context) error {
	var err error
	fmt.Printf("%s Start Check Database Version\n", infostring)
	err = checkExecuteDatabase()
	if err != nil {
		return fmt.Errorf("SQL Server Connect Error : %s", err)
	}
	fmt.Printf("%s SQL Server Connect\n", okstring)
	v, err := checkDatabaseVersion()
	if err != nil {
		return fmt.Errorf("SQL Server Version Data Error : %s", err)
	}
	fmt.Printf("%s Database Version %s\n", infostring, v.String())
	return nil
}

func updateDatabaseAction(c *cli.Context) error {
	var err error
	fmt.Printf("%s Start Update Database Version\n", infostring)
	err = checkDatabaseAction(c)
	if err != nil {
		return err
	}
	return nil
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

func checkDatabaseVersion() (version, error) {
	v := version{mejor: 0, miner: 0, revision: 0}
	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	var rows *sql.Rows
	// table check
	defer db.Close()
	rows, err = db.Query("select count(*) from sys.tables where name = N'SystemVersion' ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var count int
		rows.Scan(&count)
		if count == 0 {
			return v, nil
		}
	}

	// version check
	rows, err = db.Query("select Mejor, Miner, Revison from dbo.SystemTable where TableName = N'SystemVersion' ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var mejor, miner, revision int
		rows.Scan(
			&mejor,
			&miner,
			&revision)
		v.mejor = mejor
		v.miner = miner
		v.revision = revision
	}
	return v, nil
}

func queryVersionAction(c *cli.Context) error {
	vs, err := checkQueryVersion(folder)
	if err != nil {
		return err
	}
	for _, v := range vs {
		fmt.Println(v.String())
	}
	return nil
}

func checkQueryVersion(dir string) ([]version, error) {
	// folder layout
	//  + SystemVersion.sql
	//  + TableA
	//  | + TableA.sql
	//  | + data
	//  |   + v0.0.1.sql
	//  |   + v0.0.2.sql
	//  |   ....
	//  + TableB
	//    + TableB.sql
	//    + data
	//      + v0.0.1.sql
	//      + v0.0.2.sql
	//      ....

	versions := map[string]version{}
	ret := make([]version, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ret, err
	}

	for _, file := range files {
		if file.IsDir() {
			path := filepath.Join(dir, file.Name(), "data")
			if data, err := os.Stat(path); err == nil {
				if data.IsDir() {
					vs, err := ioutil.ReadDir(path)
					if err == nil {
						for _, v := range vs {
							filename := v.Name()
							if filename[0] == 'v' {
								// create version data for filename
								filename = filename[1:]
								temp := strings.Split(filename, ".")
								if len(temp) >= 3 {
									mejor, err1 := strconv.Atoi(temp[0])
									miner, err2 := strconv.Atoi(temp[1])
									revision, err3 := strconv.Atoi(temp[2])
									if err1 == nil && err2 == nil && err3 == nil {
										v := version{mejor: mejor, miner: miner, revision: revision}
										versions[v.String()] = v
									}
								}
							}
						}
					}
				}
			}
		}
	}

	for _, v := range versions {
		ret = append(ret, v)
	}
	sort.SliceStable(ret, func(i, j int) bool { return ret[i].compare(ret[j]) > 0 })
	return ret, nil
}
