package resolvers

import (
	"github.com/schartey/gqlgen-auth-starter/graphql"
	"github.com/schartey/gqlgen-auth-starter/user"
)

type RootResolver struct {
	userService *user.UserService
}

func NewRootResolver(userService *user.UserService) *RootResolver {
	return &RootResolver{userService}
}

func (r *RootResolver) Query() graphql.QueryResolver {
	return &rootQueryResolver{userService: r.userService}
}

func (r *RootResolver) Mutation() graphql.MutationResolver {
	return &rootMutationResolver{userService: r.userService}
}
