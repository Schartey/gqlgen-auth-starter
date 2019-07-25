package user

import (
	"time"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUsers() map[string]*User {
	return map[string]*User{
		"1": {
			ID:       "1",
			Username: "Joe",
			Person: Person{
				ID:        "1",
				Firstname: "John",
				Lastname:  "Doe",
				Email:     "john.doe@mail.com",
				Phone:     "+1234567890",
				Birthdate: time.Now(),
			},
		},
		"2": {
			ID:       "2",
			Username: "Jane",
			Person: Person{
				ID:        "2",
				Firstname: "Jane",
				Lastname:  "Doe",
				Email:     "jane.doe@mail.com",
				Phone:     "+1345678901",
				Birthdate: time.Now(),
			},
		},
	}
}
