package models

type MagicalGirl struct {
	Key            string    `json:"key"`
	Name           string    `json:"name"`
	Version        string    `json:"version"`
	IsLimited      bool      `json:"isLimited"`
	MinRare        int       `json:"minRare"`
	MaxRare        int       `json:"maxRare"`
	Attribute      Attribute `json:"attribute"`
	Type           Type      `json:"type"`
	Disk           Disk      `json:"disk"`
	Connect        Skill     `json:"connect"`
	Magia          Skill     `json:"magia"`
	Doppel         Skill     `json:"doppel"`
	Status         Status    `json:"status"`
	AwakeRate      AwakeRate `json:"awakeRate"`
	UltimetAbility Ability   `json:"ultimetAbility"`
}

type Ability struct {
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
}

type Skill struct {
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
	UseTurn int      `json:"useTurn"`
}
