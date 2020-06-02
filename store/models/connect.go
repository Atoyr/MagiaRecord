package models

type Connect struct {
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
}

type ConnectEffect struct {
	Effect   Effect
	Target   string
	TurmTurn int
}
