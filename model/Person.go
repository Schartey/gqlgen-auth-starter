package model

import "time"

type Person struct {
	ID			string
	Firstname 	string
	Lastname	string
	Email		string
	Phone		string
	Birthdate	time.Time
}