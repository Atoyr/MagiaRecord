package models

type MagicalGirl struct {
	Key            string         `json:"key"`
	Name           string         `json:"name"`
	Version        string         `json:"version"`
	IsLimited      bool           `json:"isLimited"`
	MinRare        int            `json:"minRare"`
	MaxRare        int            `json:"maxRare"`
	Attribute      Attribute      `json:"attribute"`
	Type           Type           `json:"type"`
	Disk           Disk           `json:"disk"`
	Connect        Connect        `json:"connect"`
	Magia          Magia          `json:"magia"`
	Doppel         Doppel         `json:"doppel"`
	Status         Status         `json:"status"`
	AwakeRate      AwakeRate      `json:"awakeRate"`
	UltimetAbility UltimetAbility `json:"ultimetAbility"`
}

type Magia struct {
	Name    string       `json:"name"`
	Effects []EffectDtil `json:"effects"`
}
type Doppel struct {
	Name    string       `json:"name"`
	Effects []EffectDtil `json:"effects"`
}
type UltimetAbility struct {
	Name    string       `json:"name"`
	Effects []EffectDtil `json:"effects"`
}
