//+build wireinject

package main

import (
	"context"
	"github.com/schartey/gqlgen-auth-starter/graphql/handler"
	"github.com/schartey/gqlgen-auth-starter/graphql/resolvers"
	"github.com/schartey/gqlgen-auth-starter/keycloak"
	"github.com/schartey/gqlgen-auth-starter/user"
	"github.com/spf13/viper"

	"github.com/google/wire"
)

var keycloakSet = wire.NewSet(keycloak.NewKeycloakService, keycloak.NewKeycloakRepository, keycloak.NewKeycloak)
var userSet = wire.NewSet(user.NewUserService)

func WireRootResolver(ctx context.Context, keycloakConfig *viper.Viper) *resolvers.RootResolver {
	wire.Build(resolvers.NewRootResolver, userSet, keycloakSet)
	return &resolvers.RootResolver{}
}

func WireHandlerServiceProvider(ctx context.Context, keycloakConfig *viper.Viper) *handler.HandlerServiceProvider {
	wire.Build(handler.NewHandlerServiceProvider, keycloakSet)

	return &handler.HandlerServiceProvider{}
}
