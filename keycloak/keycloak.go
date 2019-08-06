package keycloak

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var keycloakLock = &sync.Mutex{}

type KeycloakInfo struct {
	hostName     string
	clientID     string
	clientSecret string
	realm        string
	redirectURL  string
}

type OAuthInfo struct {
	provider     *oidc.Provider
	oauth2Config *oauth2.Config
	verifier     *oidc.IDTokenVerifier
}

type Keycloak struct {
	keycloakInfo *KeycloakInfo
	oAuthInfo    *OAuthInfo
}

var keycloakInstace *Keycloak

func NewKeycloak(ctx context.Context, keycloakConfig *viper.Viper) *Keycloak {

	keycloakLock.Lock()
	defer keycloakLock.Unlock()

	if keycloakInstace != nil {
		return keycloakInstace
	}

	keycloakInfo := &KeycloakInfo{
		hostName:     keycloakConfig.GetString("host-name"),
		clientID:     keycloakConfig.GetString("client-id"),
		clientSecret: keycloakConfig.GetString("client-secret"),
		realm:        keycloakConfig.GetString("realm"),
		redirectURL:  keycloakConfig.GetString("redirect-url"),
	}

	oauthInfo, err := setupOAuth(ctx, keycloakInfo)
	if err != nil {
		log.Panicf("Could not setup OAuth with Keycloak: %s", err)
	}

	keycloakInstace = &Keycloak{keycloakInfo: keycloakInfo, oAuthInfo: oauthInfo}
	return keycloakInstace
}

func connectKeycloak(ctx context.Context, hostName string) (provider *oidc.Provider, err error) {
	for i := 0; i < 10; i++ {
		log.Infof("Connecting to Keycloak")
		provider, perr := oidc.NewProvider(ctx, hostName)

		if perr == nil {
			log.Infof("Connected to Keycloak")
			return provider, nil
		}
		err = perr
		log.Errorf("%s", err.Error())
		time.Sleep(10 * time.Second)
	}
	return nil, err
}

func setupOAuth(ctx context.Context, keycloakInfo *KeycloakInfo) (*OAuthInfo, error) {
	provider, err := connectKeycloak(ctx, keycloakInfo.hostName)
	if err != nil {
		return nil, err
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := &oauth2.Config{
		ClientID:     keycloakInfo.clientID,
		ClientSecret: keycloakInfo.clientSecret,
		RedirectURL:  keycloakInfo.redirectURL,
		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}

	oidcConfig := &oidc.Config{
		ClientID: keycloakInfo.clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	oauthInfo := &OAuthInfo{
		provider:     provider,
		oauth2Config: oauth2Config,
		verifier:     verifier,
	}

	return oauthInfo, nil
}
