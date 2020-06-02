package models

type Effect struct {
	Key                  string `json:"key"`
	Name                 string `json:"name"`
	EffectTypeName       string `json:"effectType"`
	EffectCategoryName   string `json:"effectCategory"`
	EffectActionName     string `json:"effectAction"`
	EffectActivationName string `json:"effectActivation"`
}

type EffectDtil struct {
	*Effect
	Target string `json:"target"`
	Term   int    `json:"term"`
}
