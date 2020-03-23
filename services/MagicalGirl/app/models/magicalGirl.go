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

	// Get Attribute
	ctx = context.Background()
	attrrows, err := db.QueryContext(ctx, selectAttributeQuery)
	if err != nil {
		return nil, err
	}
	defer attrrows.Close()
	attributes := make(map[int]string, 0)
	for attrrows.Next() {
		var id int
		var name string
		if err = attrrows.Scan(
			&id,
			&name,
		); err != nil {
			return nil, err
		}
		attributes[id] = name
	}

	// Get Type
	ctx = context.Background()
	typerows, err := db.QueryContext(ctx, selectTypeQuery)
	if err != nil {
		return nil, err
	}
	defer typerows.Close()
	types := make(map[int]string, 0)
	for typerows.Next() {
		var id int
		var name string
		if err = typerows.Scan(
			&id,
			&name,
		); err != nil {
			return nil, err
		}
		types[id] = name
	}

	// Get Magical Girl
	ctx = context.Background()
	mgrows, err := db.QueryContext(ctx, selectMagicalGirlQuery)
	if err != nil {
		return nil, err
	}
	defer mgrows.Close()

	magicalGirls := make([]MagicalGirl, 0)
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
		mg.Attribute = attributes[attributeID]
		mg.Type = types[typeID]

		magicalGirls = append(magicalGirls, mg)
	}

	return magicalGirls, nil
}
