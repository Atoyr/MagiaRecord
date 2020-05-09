package models

import (
	"database/sql"
)

type Attribute struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetAttributes(db *sql.DB) ([]Attribute, error) {
	as := make([]Attribute, 0)
	rows, err := db.Query(`
  select 
    ID
    ,Name
  from Attributes
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a Attribute
		rows.Scan(&a.Id, &a.Name)
		as = append(as, a)
	}
	return as, nil
}
