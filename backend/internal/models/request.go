package models

type Request struct {
	RequestId   uint64
	UserId      uint64
	UserLogin   string
	Description string
	Status      string
}
