package models

import (
	"context"
	"database/sql"

	"github.com/atoyr/MagiaRecord/services/MagicalGirl/app"
	_ "github.com/denisenkom/go-mssqldb"
)

const (
	fire  = "FIRE"
	water = "WATER"
	tree  = "TREE"
	light = "LIGHT"
	dark  = "DARK"
	none  = "NONE"
)

// MagicalGirl Models
type MagicalGirl struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	RomanName string `json:"roman_name"`
	Attribute string `json:"attribute"`
	Type      string `json:"type"`
}

func GetMagicalGirlAll() ([]MagicalGirl, error) {
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

	ctx = context.Background()
	rows, err := db.QueryContext(ctx, "")
	if err != nil {
		return nil, err
	}

	magicalGirls := make([]MagicalGirl, 0)
	defer rows.Close()
	for rows.Next() {
		mg := MagicalGirl{}
		if err = rows.Scan(
			&mg.ID,
			&mg.Name,
			&mg.Version,
			&mg.RomanName,
			&mg.Attribute,
			&mg.Type,
		); err != nil {
			return nil, err
		}

		magicalGirls = append(magicalGirls, mg)
	}

	return magicalGirls, nil
}
