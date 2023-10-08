package models

type Farm struct {
	FarmId      uint64
	Name        string
	Description string
	Address     string
	Owner       string
	Honey       []Honey
}

type FarmPatch struct {
	Description string
	Address     string
	Honey       []Honey
}
