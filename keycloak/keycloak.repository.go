package keycloak

import "github.com/Nerzal/gocloak/v3"

type KeycloakRepository struct {
	// All this info and setup should probably go in some kind of connection manager (like db connection)
	keycloak *Keycloak
	rest *KeycloakRest

	roleMapping		map[string]*gocloak.Role
}

func NewKeycloakRepository(keycloak *Keycloak, rest *KeycloakRest) *KeycloakRepository {
	roleMapping, err := rest.GetClientRoleMapping()
	if err != nil {
		panic(err)
	}
	return &KeycloakRepository{keycloak: keycloak, rest: rest, roleMapping: roleMapping}
}

func (k *KeycloakRepository) GetRoleMappingsByUserId(userId string) ([]*UserRole, error) {
	roleMappings, err := k.rest.GetRoleMappingsByUserId(userId)
	if (err != nil) {
		return nil, err
	}

	var userRoles []*UserRole

	for _, element := range roleMappings {
		userRoles = append(userRoles, &UserRole{Name: element.Name})
	}

	return userRoles, nil
}