package models

type Request struct {
	RequestId   uint64
	UserLogin   string
	Description string
	Status     string
}

type AllRequests struct {
	Requests []Request
}
