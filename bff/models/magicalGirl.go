package models

type MagicalGirl struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Attribute string `json:"attribute"`
	Disk      Disk   `json:"disk"`
}

func GetMagicalGirls() []MagicalGirl {
	return nil
}
