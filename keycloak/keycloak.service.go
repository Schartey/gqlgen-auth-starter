package keycloak

import (
	"context"
	"github.com/coreos/go-oidc"
)

type KeycloakService struct {
	keycloakRepository *KeycloakRepository
}

func NewKeycloakService(keycloakRepository *KeycloakRepository) *KeycloakService {
	return &KeycloakService{keycloakRepository: keycloakRepository}
}

func (k *KeycloakService) GetAuthCodeURL() string {
	state := "test"
	return k.keycloakRepository.keycloak.oAuthInfo.oauth2Config.AuthCodeURL(state)
}

func (k *KeycloakService) Verify(ctx context.Context, part string) (token *oidc.IDToken, err error) {
	return k.keycloakRepository.keycloak.oAuthInfo.verifier.Verify(ctx, part)
}
