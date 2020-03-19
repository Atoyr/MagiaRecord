package models

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
	ID        string `json:"id"`
	Name      string `json:"name"`
	KanaName  string `json:"kana_name"`
	Version   string `json:"version"`
	Attribute string `json:"attribute"`
	Type      string `json:"type"`
	Disk      string `json:"disk"`
}
