package models

type Request struct {
	RequestId     uint64
	UserLogin     string
	Description   string
	State         string
	AdminResponse string
}

type AllRequests struct {
	Requests []Request
}
