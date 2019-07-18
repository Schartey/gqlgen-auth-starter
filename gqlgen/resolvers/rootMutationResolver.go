package resolvers

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/gqlgen"
	"github.com/schartey/gqlgen-auth-starter/model"
	"strconv"
)

type rootMutationResolver struct {
	*Resolver
	Users map[string]*model.User
}

func (r *rootMutationResolver) AddUser(ctx context.Context, user gqlgen.UserInput) (*model.User, error) {
	id := 0
	for k := range r.Users {
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
	r.Users[idString] = &model.User{
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

	return r.Users[idString], nil
}
