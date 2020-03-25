package models

import (
	"context"
	"database/sql"

	"github.com/atoyr/MagiaRecord/services/Attribute/app"
	_ "github.com/denisenkom/go-mssqldb"
)

// Attribute Models
type Attribute struct {
	ID         int    `json:"id"`
	SystemName string `json:"system_name"`
	Name       string `json:"name"`
}

func GetAttributeAll() ([]Attribute, error) {
	c := app.NewConfig()
	var err error

	db, err := sql.Open("sqlserver", c.ConnString())
	if err != nil {

	}
	var ctx context.Context
	ctx = context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Get Magical Girl
	ctx = context.Background()
	rows, err := db.QueryContext(ctx, selectAttributeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	attributes := make([]Attribute, 0)
	for rows.Next() {
		a := Attribute{}
		if err = rows.Scan(
			&a.ID,
			&a.Name,
		); err != nil {
			return nil, err
		}

		attributes = append(attributes, a)
	}

	return attributes, nil
}
