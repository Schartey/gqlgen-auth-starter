//+build wireinject

package main

import (
	"github.com/schartey/gqlgen-auth-starter/graphql/resolvers"
	"github.com/schartey/gqlgen-auth-starter/user"

	"github.com/google/wire"
)

func WireUp() *resolvers.RootResolver {
	wire.Build(resolvers.NewRootResolver, user.NewUserService)
	return &resolvers.RootResolver{}
}
