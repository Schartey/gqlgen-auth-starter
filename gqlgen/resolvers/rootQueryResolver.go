package resolvers

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

type rootQueryResolver struct {
	*Resolver
	Users map[string]*model.User
}

func (r *rootQueryResolver) GetUserByID(ctx context.Context, id *string) (*model.User, error) {
	user := r.Users[*id]
	return user, nil
}

func (r *rootQueryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	for _, user := range r.Users {
		result = append(result, user)
	}
	return result, nil
}
