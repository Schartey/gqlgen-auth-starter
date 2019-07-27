//+build wireinject

package main

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/graphql/resolvers"
	"github.com/schartey/gqlgen-auth-starter/keycloak"
	"github.com/schartey/gqlgen-auth-starter/user"
	"github.com/spf13/viper"

	"github.com/google/wire"
)

func WireUp(ctx context.Context, keycloakConfig *viper.Viper) *resolvers.RootResolver {
	wire.Build(resolvers.NewRootResolver, user.NewUserService, keycloak.NewKeycloakService, keycloak.NewKeycloakRepository, keycloak.NewKeycloak)
	return &resolvers.RootResolver{}
}
