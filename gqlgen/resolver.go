package gqlgen

import (
	"context"

	"github.com/schartey/gqlgen-auth-starter/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{
	Users map[string]*model.User
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUserByID(ctx context.Context, id *string) (*model.User, error) {
	user := r.Users[*id]
	return user, nil
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	for _, user := range r.Users {
		result = append(result, user)
	}
	return result, nil
}

