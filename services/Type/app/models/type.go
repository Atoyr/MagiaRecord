package models

import (
	"context"
	"database/sql"

	"github.com/atoyr/MagiaRecord/services/Type/app"
	_ "github.com/denisenkom/go-mssqldb"
)

// Type Models
type Type struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetTypeAll() ([]Type, error) {
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
	rows, err := db.QueryContext(ctx, selectTypeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	types := make([]Type, 0)
	for rows.Next() {
		t := Type{}
		if err = rows.Scan(
			&t.ID,
			&t.Name,
		); err != nil {
			return nil, err
		}

		types = append(types, t)
	}

	return types, nil
}
