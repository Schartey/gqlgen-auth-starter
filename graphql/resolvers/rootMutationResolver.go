package resolvers

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/graphql"
	"github.com/schartey/gqlgen-auth-starter/user"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type rootMutationResolver struct {
	*Resolver
	userService *user.UserService
}

func (r *rootMutationResolver) AddUser(ctx context.Context, userInput graphql.UserInput) (*user.User, error) {
	id := 0
	for k := range r.userService.GetUsers() {
		currentId, err := strconv.Atoi(k)
		if err != nil {
			return nil, err
		}
		if id < currentId {
			id = currentId
		}
	}
	id++
	idString := strconv.Itoa(id)
	r.userService.GetUsers()[idString] = &user.User{
		Username: userInput.Username,
		Person: user.Person{
			Firstname: userInput.Person.Firstname,
			Lastname:  userInput.Person.Lastname,
			Email:     userInput.Person.Email,
			Phone:     userInput.Person.Phone,
			Birthdate: userInput.Person.Birthdate,
			ID:        "shouldbegenerated",
		},
	}

	log.WithField("user", userInput).Debugf("AddUser")

	return r.userService.GetUsers()[idString], nil
}
