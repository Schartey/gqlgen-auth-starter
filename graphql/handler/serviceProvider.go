package handler

import (
	"github.com/schartey/gqlgen-auth-starter/keycloak"
	"github.com/schartey/gqlgen-auth-starter/user"
)

type HandlerServiceProvider struct {
	KeycloakService *keycloak.KeycloakService
	UserService		*user.UserService
}

func NewHandlerServiceProvider(userService *user.UserService, keycloakService *keycloak.KeycloakService) *HandlerServiceProvider {
	return &HandlerServiceProvider{KeycloakService: keycloakService, UserService: userService}
}
