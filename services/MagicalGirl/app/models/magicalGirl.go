package models

import (
	"context"
	"database/sql"

	"github.com/atoyr/MagiaRecord/services/MagicalGirl/app"
	_ "github.com/denisenkom/go-mssqldb"
)

// MagicalGirl Models
type MagicalGirl struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	RomanName   string `json:"roman_name"`
	AttributeID int    `json:"attribute_id"`
	TypeID      int    `json:"type_id"`
	DiskID      int    `json:"disk_id"`
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
		if err = mgrows.Scan(
			&mg.ID,
			&mg.Name,
			&mg.Version,
			&mg.RomanName,
			&mg.AttributeID,
			&mg.TypeID,
			&mg.DiskID,
		); err != nil {
			return nil, err
		}

		magicalGirls = append(magicalGirls, mg)
	}

	return magicalGirls, nil
}
