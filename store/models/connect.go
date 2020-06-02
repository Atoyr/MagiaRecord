package models

type Connect struct {
	Name    string          `json:"name"`
	Effects []ConnectEffect `json:"effects"`
}

type ConnectEffect struct {
	Effect Effect
	Target string
	Turm   int
}
