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

type Attribute struct {
	ID   int
	Name string
}

type Type struct {
	ID   int
	Name string
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
	mgrows, err := db.QueryContext(ctx, "")
	if err != nil {
		return nil, err
	}

	magicalGirls := make([]MagicalGirl, 0)
	defer mgrows.Close()
	for mgrows.Next() {
		mg := MagicalGirl{}
		var attributeID, typeID int
		if err = mgrows.Scan(
			&mg.ID,
			&mg.Name,
			&mg.Version,
			&mg.RomanName,
			&attributeID,
			&typeID,
		); err != nil {
			return nil, err
		}

		magicalGirls = append(magicalGirls, mg)
	}

	return magicalGirls, nil
}
