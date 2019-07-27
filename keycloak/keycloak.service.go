package keycloak

type KeycloakService struct {
	// All this info and setup should probably go in some kind of connection manager (like db connection)
	keycloakRepository *KeycloakRepository
}

func NewKeycloakService(keycloakRepository *KeycloakRepository) *KeycloakService {
	return &KeycloakService{keycloakRepository: keycloakRepository}
}
