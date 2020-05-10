package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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
var infostring = "[\x1b[36mINFO\x1b[0m]"

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
			Action: getDatabaseVersionAction,
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
			Action: getQueryVersionAction,
		},
		{
			Name:    "QueryTableToVersion",
			Aliases: []string{"qtv"},
			Usage:   "Get Query Table to Version",
			Flags: []cli.Flag{
				folderFlag,
			},
			Action: getQueryTable2VersionAction,
		},
		{
			Name:    "QueryVersionToTable",
			Aliases: []string{"qvt"},
			Usage:   "Get Query Version to Table",
			Flags: []cli.Flag{
				folderFlag,
			},
			Action: getQueryVersion2TableAction,
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

func getDatabaseVersionAction(c *cli.Context) error {
	var err error
	fmt.Printf("%s Start Check Database Version\n", infostring)

	err = tryDatabaseConnect()
	if err != nil {
		return fmt.Errorf("SQL Server Connect Error : %s", err)
	}
	fmt.Printf("%s SQL Server Connect\n", okstring)

	ok, err := isExistsSystemVersionTable()
	if err != nil {
		return fmt.Errorf("SystemVersion Check error : %s", err)
	}
	if ok {
		v, err := getSystemVersion()
		if err != nil {
			return fmt.Errorf("SystemVersion Check error : %s", err)
		}
		fmt.Printf("%s Database Version %s\n", infostring, v.String())
	} else {
		fmt.Printf("%s SystemVersion Table is not exists\n", infostring)
	}

	return nil
}

func updateDatabaseAction(c *cli.Context) error {
	var err error
	fmt.Printf("%s Start Update Database Version\n", infostring)
	err = getDatabaseVersionAction(c)
	if err != nil {
		return err
	}
	return nil
}

func getQueryVersionAction(c *cli.Context) error {
	vs, err := getVersionAtQuery(folder)
	if err != nil {
		return err
	}
	for _, v := range vs {
		fmt.Println(v.String())
	}
	return nil
}

func getQueryTable2VersionAction(c *cli.Context) error {
	tables, err := getTable2VersionAtQuery(folder)
	if err != nil {
		return err
	}
	for k, vs := range tables {
		fmt.Println(k)
		for _, v := range vs {
			fmt.Printf("  %s\n", v.String())
		}
	}
	return nil
}

func getQueryVersion2TableAction(c *cli.Context) error {
	vs, err := getVersion2TableAtQuery(folder)
	if err != nil {
		return err
	}
	for v, ts := range vs {
		fmt.Println(v)
		for _, t := range ts {
			fmt.Printf("  %s\n", t)
		}
	}
	return nil
}

func tryDatabaseConnect() error {
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

func isExistsSystemVersionTable() (bool, error) {
	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		return false, err
	}
	defer db.Close()

	rows, err := db.Query("select count(*) from sys.tables where name = N'SystemVersion' ")
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		var count int
		rows.Scan(&count)
		if count == 0 {
			return false, nil
		}
	}

	return true, nil
}

func getSystemVersion() (version, error) {
	v := version{mejor: 0, miner: 0, revision: 0}
	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		return v, err
	}
	defer db.Close()

	rows, err := db.Query("select Mejor, Miner, Revison from dbo.SystemVersion where TableName = N'SystemVersion' ")
	if err != nil {
		return v, err
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

func getTableVersion() (map[string]version, error) {
	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tableVersion := map[string]version{}
	rows, err := db.Query("select TableName, Mejor, Miner, Revison from dbo.SystemTable where TableName <> N'SystemVersion' ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tableName string
		var mejor, miner, revision int
		rows.Scan(
			&tableName,
			&mejor,
			&miner,
			&revision)
		v := version{}
		v.mejor = mejor
		v.miner = miner
		v.revision = revision
		tableVersion[tableName] = v
	}

	return tableVersion, nil
}

func createTable(dir, tableName string) error {
	path := filepath.Join(dir, tableName, fmt.Sprintf("%s.sql", tableName))
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(string(b))
	if err != nil {
		return err
	}

	return nil
}

func executeVersionData(dir, tableName string, v version) error {
	path := filepath.Join(dir, tableName, fmt.Sprintf("v%s.sql", v.String()))
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlserver", getConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(string(b))
	if err != nil {
		return err
	}

	return nil
}

// getTable2VersionAtQuery is return map[TableName string][]version for DB Query folder
func getTable2VersionAtQuery(dir string) (map[string][]version, error) {
	tablefolders, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	ret := map[string][]version{}
	for _, tablefolder := range tablefolders {
		if tablefolder.IsDir() {

			tableQueryPath := filepath.Join(dir, tablefolder.Name(), fmt.Sprintf("%s.sql", tablefolder.Name()))
			if _, err := os.Stat(tableQueryPath); err == nil {
				versions := make([]version, 0)

				datafolderPath := filepath.Join(dir, tablefolder.Name(), "data")
				if data, err := os.Stat(datafolderPath); err == nil {
					if data.IsDir() {
						versionfiles, err := ioutil.ReadDir(datafolderPath)
						if err == nil {
							for _, versionfile := range versionfiles {
								filename := versionfile.Name()
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
											versions = append(versions, v)
										}
									}
								}
							}
						}
					}
				}

				sort.SliceStable(versions, func(i, j int) bool { return versions[i].compare(versions[j]) > 0 })
				ret[tablefolder.Name()] = versions
			}
		}
	}
	return ret, nil
}

func getVersion2TableAtQuery(dir string) (map[string][]string, error) {
	tableVersions, err := getTable2VersionAtQuery(dir)
	if err != nil {
		return nil, err
	}
	ret := map[string][]string{}
	for k, vs := range tableVersions {
		for _, v := range vs {
			if _, ok := ret[v.String()]; ok {
				ret[v.String()] = append(ret[v.String()], k)
			} else {
				ret[v.String()] = []string{k}
			}
		}
	}
	return ret, nil
}

func getVersionAtQuery(dir string) ([]version, error) {
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
	// return 0.0.1 and 0.0.2

	versions := map[string]version{}
	ret := make([]version, 0)
	tableVersions, err := getTable2VersionAtQuery(dir)
	if err != nil {
		return nil, err
	}

	for _, vs := range tableVersions {
		for _, v := range vs {
			versions[v.String()] = v
		}
	}
	for _, v := range versions {
		ret = append(ret, v)
	}
	sort.SliceStable(ret, func(i, j int) bool { return ret[i].compare(ret[j]) > 0 })
	return ret, nil
}
