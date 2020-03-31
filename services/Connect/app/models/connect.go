package models

import (
	"context"
	"database/sql"

	"github.com/atoyr/MagiaRecord/services/Connect/app"
	_ "github.com/denisenkom/go-mssqldb"
)

// Connect Models
type Connect struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	MagicalGirlID int    `json:"magical_girl_id"`
}

func GetConnectAll() ([]Connect, error) {
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

	// Get Connect
	ctx = context.Background()
	rows, err := db.QueryContext(ctx, selectConnectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	connects := make([]Connect, 0)
	for rows.Next() {
		c := Connect{}
		if err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.MagicalGirlID,
		); err != nil {
			return nil, err
		}

		connects = append(connects, c)
	}

	return connects, nil
}
