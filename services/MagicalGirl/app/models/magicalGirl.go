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
	ID   string
	Name string
	Type string
}
