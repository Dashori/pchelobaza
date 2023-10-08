package models

type Conference struct {
	ConderenceId   uint64
	Name           string
	Owner          string
	Description    string
	Date           string
	Address        string
	AvailableUsers float64
	MaxUsers       float64
}

type ConferencePatch struct {
	ConderenceId   uint64
	Description    string
	Date           string
	Address        string
	AvailableUsers float64
	MaxUsers       float64
}
