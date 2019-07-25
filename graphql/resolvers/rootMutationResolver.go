package resolvers

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/graphql"
	"github.com/schartey/gqlgen-auth-starter/model"
	"github.com/schartey/gqlgen-auth-starter/user"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type rootMutationResolver struct {
	*Resolver
	userService *user.UserService
}

func (r *rootMutationResolver) AddUser(ctx context.Context, user graphql.UserInput) (*model.User, error) {
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
	r.userService.GetUsers()[idString] = &model.User{
		Username: user.Username,
		Person: model.Person{
			Firstname: user.Person.Firstname,
			Lastname:  user.Person.Lastname,
			Email:     user.Person.Email,
			Phone:     user.Person.Phone,
			Birthdate: user.Person.Birthdate,
			ID:        "shouldbegenerated",
		},
	}

	log.WithField("user", user).Debugf("AddUser")

	return r.userService.GetUsers()[idString], nil
}
