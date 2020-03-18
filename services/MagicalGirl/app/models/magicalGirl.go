package models

const (
	fire  = "FIRE"
	water = "WATER"
	tree  = "TREE"
	light = "LIGHT"
	dark  = "DARK"
	none  = "NONE"
)

type MagicalGirl struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	KanaName string `json:"kana_name"`
	Type     string `json:"type"`
}
