package models

type ConnectSkill struct {
  ID int `json:"id"`
  ConnectID int `json:"connect_id"`
  SkillID int  `json:"skill_id"`
  Level int `json:"level"`
  EffectTurn int `json:"effect_turn"`
  Target string `json:"target"`
}


