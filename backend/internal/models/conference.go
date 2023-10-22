package models

type Conference struct {
	ConderenceId   uint64
	Name           string
	Owner          string
	Description    string
	Date           string
	Address        string
	MaxUsers       float64
	CurrentUsers float64
}

type ConferencePatch struct {
	ConderenceId   uint64
	Description    string
	Date           string
	Address        string
	MaxUsers       float64
	CurrentUsers float64
}
