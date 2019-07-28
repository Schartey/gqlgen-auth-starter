package handler

import "github.com/schartey/gqlgen-auth-starter/keycloak"

type HandlerServiceProvider struct {
	KeycloakService *keycloak.KeycloakService
}

func NewHandlerServiceProvider(keycloakService *keycloak.KeycloakService) *HandlerServiceProvider {
	return &HandlerServiceProvider{KeycloakService: keycloakService}
}
