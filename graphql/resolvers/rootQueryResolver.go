package resolvers

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/user"
	log "github.com/sirupsen/logrus"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

type rootQueryResolver struct {
	*Resolver
	userService *user.UserService
}

func (r *rootQueryResolver) GetUserByID(ctx context.Context, id *string) (*user.User, error) {
	foundUser := r.userService.GetUsers()[*id]
	log.WithField("user", foundUser).Debugf("GetUser")

	return foundUser, nil
}

func (r *rootQueryResolver) GetUsers(ctx context.Context) ([]*user.User, error) {
	var result []*user.User
	for _, currentUser := range r.userService.GetUsers() {
		result = append(result, currentUser)
	}
	log.WithField("users", result).Debugf("GetUsers")

	return result, nil
}
