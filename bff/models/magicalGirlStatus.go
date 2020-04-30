package models

type MagicalGirlStatus struct {
	*MagicalGirl
	Rare    int
	HP      int
	Attack  int
	Defence int
}
