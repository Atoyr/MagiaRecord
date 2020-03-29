package models

import (
	"context"
	"database/sql"

	"github.com/atoyr/MagiaRecord/services/Disk/app"
	_ "github.com/denisenkom/go-mssqldb"
)

// Disk Models
type Disk struct {
	ID     int  `json:"id"`
	Accele uint `json:"accele"`
	BlastH uint `json:"blast_h"`
	BlastV uint `json:"blast_v"`
	Charge uint `json:"charge"`
}

func GetDiskAll() ([]Disk, error) {
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

	// Get Disk
	ctx = context.Background()
	rows, err := db.QueryContext(ctx, selectDiskQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	disks := make([]Disk, 0)
	for rows.Next() {
		d := Disk{}
		if err = rows.Scan(
			&d.ID,
			&d.Accele,
			&d.BlastH,
			&d.BlastV,
			&d.Charge,
		); err != nil {
			return nil, err
		}

		disks = append(disks, d)
	}

	return disks, nil
}
