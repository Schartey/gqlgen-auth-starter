package keycloak

type KeycloakRepository struct {
	// All this info and setup should probably go in some kind of connection manager (like db connection)
	keycloak *Keycloak
}

func NewKeycloakRepository(keycloak *Keycloak) *KeycloakRepository {
	return &KeycloakRepository{keycloak: keycloak}
}
