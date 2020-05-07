package models

import (
	"database/sql"
)

type MagicalGirl struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Type      string `json:"type"`
	Attribute string `json:"attribute"`
	Disk      Disk   `json:"disk"`
}

func GetMagicalGirls(db *sql.DB) ([]MagicalGirl, error) {
	mgs := make([]MagicalGirl, 0)
	rows, err := db.Query(`
  select 
    MagicalGirls.ID
    ,MagicalGirls.Name
    ,MagicalGirls.Version
    ,Type.Name
    ,Attribute.Name
    ,Disk.Accele
    ,Disk.BlastV
    ,Disk.BlastH
    ,Disk.Charge
  from 
    MagicalGirls
  inner join 
    Type
  on 
    MagicalGirls.TypeID = Type.ID
  inner join
    Disk
  on
    MagicalGirls.DiskID = Disk.ID
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		mg := MagicalGirl{}
		mg.Disk = Disk{}
		rows.Scan(
			&mg.Id,
			&mg.Name,
			&mg.Version,
			&mg.Type,
			&mg.Attribute,
			&mg.Disk.Accele,
			&mg.Disk.BlastV,
			&mg.Disk.BlastH,
			&mg.Disk.Charge,
		)
		mgs = append(mgs, mg)
	}
	return mgs, nil
}
