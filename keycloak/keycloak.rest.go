package keycloak

import (
	"context"
	"errors"
	"github.com/Nerzal/gocloak/v3"
	"github.com/spf13/viper"
	"sync"
)

var keycloakRestLock = &sync.Mutex{}

type KeycloakRest struct {
	client		gocloak.GoCloak
	adminToken	*gocloak.JWT
	uid			string
	clientId	string
	realm		string
}

var keycloakRestInstance *KeycloakRest

func NewKeycloakRest(ctx context.Context, config *viper.Viper) (*KeycloakRest, error) {

	keycloakRestLock.Lock()
	defer keycloakRestLock.Unlock()

	if keycloakRestInstance != nil {
		return keycloakRestInstance, nil
	}

	hostName := config.GetString("host-name")
	clientId := config.GetString("client-id")
	clientSecret := config.GetString("client-secret")
	realm := config.GetString("realm")

	client := gocloak.NewClient(hostName)
	adminToken, err := client.LoginClient(clientId, clientSecret, realm)
	if err != nil {
		panic(err)
	}

	uid, err := getClientUID(&client, adminToken, clientId, realm)
	if err != nil {
		panic(err)
	}

	keycloakRestInstance := &KeycloakRest{client: client, adminToken: adminToken, uid: *uid, clientId: clientId, realm: realm}
	return keycloakRestInstance, nil
}

func getClientUID(client *gocloak.GoCloak, adminToken *gocloak.JWT, clientId string, realm string) (*string, error) {
	getClientsParams := &gocloak.GetClientsParams{}

	clients, err := (*client).GetClients(adminToken.AccessToken, realm, *getClientsParams)
	if err != nil {
		return nil, err
	}

	for _, element := range clients {
		if element.ClientID == clientId {
			return &element.ID, nil
		}
	}

	return nil, errors.New("Client is not registered in keycloak")
}

func (k *KeycloakRest) GetClientRoleMapping() (map[string]*gocloak.Role, error) {

	roles, err := k.client.GetClientRoles(k.adminToken.AccessToken, k.realm, k.uid)
	if err != nil {
		return nil, err
	}

	roleMapping := make(map[string]*gocloak.Role)

	for _, element := range roles {
		roleMapping[element.Name] = element
	}

	return roleMapping, nil
}

func (k *KeycloakRest) GetRoleMappingsByUserId(userId string) ([]gocloak.Role, error) {
	mappingRepresentation, err := k.client.GetRoleMappingByUserID(k.adminToken.AccessToken, k.realm, userId)
	if err != nil {
		return nil, err
	}

	return mappingRepresentation.ClientMappings[k.clientId].Mappings, nil
}