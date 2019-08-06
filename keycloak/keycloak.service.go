package keycloak

import (
	"context"
	"encoding/json"
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

func (k *KeycloakService) VerifyCallback(ctx context.Context, code string) (interface{}, error) {
	oauth2Token, err := k.keycloakRepository.keycloak.oAuthInfo.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return nil, err
	}

	return rawIDToken, nil
}

func (k *KeycloakService) ConvertToken(ctx context.Context, rawIDToken interface{}) (*User, error) {
	//Split up into calls
	idToken, err := k.keycloakRepository.keycloak.oAuthInfo.verifier.Verify(ctx, rawIDToken.(string))
	if err != nil {
		return nil, err
	}

	var IDTokenClaims *json.RawMessage

	if err := idToken.Claims(&IDTokenClaims); err != nil {
		return nil, err
	}

	var userToken UserToken
	err = json.Unmarshal(*IDTokenClaims, &userToken)
	if err != nil {
		return nil, err
	}

	userRoles, err := k.GetRolesByUserId(userToken.Sub)
	if err != nil {
		return nil, err
	}

	user := &User{
		UserToken: userToken,
		RawIDToken: rawIDToken,
		UserRoles: userRoles,
	}

	return user, nil
}

func (k *KeycloakService) GetRolesByUserId(userId string) ([]*UserRole, error) {
	return k.keycloakRepository.GetRoleMappingsByUserId(userId)
}
