package models

import "time"

type User struct {
	UserId           uint64
	Login            string
	Password         string
	ConfirmPassword  string
	Name             string
	Surname          string
	Contacts         string
	RegistrationDate time.Time
	Role             string
}

type UserPatch struct {
	Password string
	Name     string
	Surname  string
	Contacts string
}

// type NewUser struct {
// 	Login           string
// 	Password        string
// 	ConfirmPassword string
// 	Name            string
// 	Surname         string
// 	Contacts        string
// }
