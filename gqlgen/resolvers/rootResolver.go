package resolvers

import (
	"github.com/schartey/gqlgen-auth-starter/gqlgen"
	"github.com/schartey/gqlgen-auth-starter/model"
)

type RootResolver struct {
	Users map[string]*model.User
}

func NewRootResolver(users map[string]*model.User) *RootResolver {
	return &RootResolver{users}
}

func (r *RootResolver) Query() gqlgen.QueryResolver {
	return &rootQueryResolver{Users: r.Users}
}

func (r *RootResolver) Mutation() gqlgen.MutationResolver {
	return &rootMutationResolver{Users: r.Users}
}
