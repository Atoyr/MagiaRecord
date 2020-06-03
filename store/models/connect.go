package models

type Connect struct {
	Name    string       `json:"name"`
	Effects []EffectDtil `json:"effects"`
}
