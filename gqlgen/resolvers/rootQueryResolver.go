package resolvers

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/model"
	log "github.com/sirupsen/logrus"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

type rootQueryResolver struct {
	*Resolver
	Users map[string]*model.User
}

func (r *rootQueryResolver) GetUserByID(ctx context.Context, id *string) (*model.User, error) {
	user := r.Users[*id]
	log.WithField("user", user).Debugf("GetUser")

	return user, nil
}

func (r *rootQueryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	for _, user := range r.Users {
		result = append(result, user)
	}
	log.WithField("users", result).Debugf("GetUsers")

	return result, nil
}
