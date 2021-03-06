// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/schartey/gqlgen-auth-starter/graphql/handler"
	"github.com/schartey/gqlgen-auth-starter/graphql/resolvers"
	"github.com/schartey/gqlgen-auth-starter/keycloak"
	"github.com/schartey/gqlgen-auth-starter/user"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

func WireRootResolver(ctx context.Context, keycloakConfig *viper.Viper) *resolvers.RootResolver {
	keycloakKeycloak := keycloak.NewKeycloak(ctx, keycloakConfig)
	keycloakRepository := keycloak.NewKeycloakRepository(keycloakKeycloak)
	keycloakService := keycloak.NewKeycloakService(keycloakRepository)
	userService := user.NewUserService(keycloakService)
	rootResolver := resolvers.NewRootResolver(userService)
	return rootResolver
}

func WireHandlerServiceProvider(ctx context.Context, keycloakConfig *viper.Viper) *handler.HandlerServiceProvider {
	keycloakKeycloak := keycloak.NewKeycloak(ctx, keycloakConfig)
	keycloakRepository := keycloak.NewKeycloakRepository(keycloakKeycloak)
	keycloakService := keycloak.NewKeycloakService(keycloakRepository)
	handlerServiceProvider := handler.NewHandlerServiceProvider(keycloakService)
	return handlerServiceProvider
}

// wire.go:

var keycloakSet = wire.NewSet(keycloak.NewKeycloakService, keycloak.NewKeycloakRepository, keycloak.NewKeycloak)

var userSet = wire.NewSet(user.NewUserService)
