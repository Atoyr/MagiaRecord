package models

import (
	"database/sql"
)

type Type struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetTypes(db *sql.DB) ([]Type, error) {
	ts := make([]Type, 0)
	rows, err := db.Query(`
  select 
    ID
    ,Name
  from Types
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t Type
		rows.Scan(&t.Id, &t.Name)
		ts = append(ts, t)
	}
	return ts, nil
}
