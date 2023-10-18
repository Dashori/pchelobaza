package models

import "time"

type User struct {
	UserId          uint64
	Login           string
	Password        string
	ConfirmPassword string
	Name            string
	Surname         string
	Contact         string
	RegisteredAt    time.Time
	Role            string
}

type UserPatch struct {
	Login    string
	Password string
	Name     string
	Surname  string
	Contact  string
}

// type NewUser struct {
// 	Login           string
// 	Password        string
// 	ConfirmPassword string
// 	Name            string
// 	Surname         string
// 	Contacts        string
// }
