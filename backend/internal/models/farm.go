package models

type Farm struct {
	FarmId      uint64
	UserId      uint64
	UserLogin   string
	Name        string
	Description string
	Address     string
	Honey       []Honey
}
