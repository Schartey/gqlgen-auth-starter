package user

import (
	"context"
	"errors"
	"github.com/schartey/gqlgen-auth-starter/keycloak"
	"time"
)

type UserService struct {
	keycloakService *keycloak.KeycloakService
}

func NewUserService(keycloakService *keycloak.KeycloakService) *UserService {
	return &UserService{keycloakService: keycloakService}
}

func (u *UserService) GetUsers() map[string]*User {
	return map[string]*User{
		"1": {
			ID:       "1",
			Username: "Joe",
			Person: Person{
				ID:        "1",
				Firstname: "John",
				Lastname:  "Doe",
				Email:     "john.doe@mail.com",
				Phone:     "+1234567890",
				Birthdate: time.Now(),
			},
		},
		"2": {
			ID:       "2",
			Username: "Jane",
			Person: Person{
				ID:        "2",
				Firstname: "Jane",
				Lastname:  "Doe",
				Email:     "jane.doe@mail.com",
				Phone:     "+1345678901",
				Birthdate: time.Now(),
			},
		},
	}
}

func (u *UserService) HandleLogin(ctx context.Context, state string, code string) (*keycloak.User, error) {
	//State must be random!
	if "test" != state {
		return nil, errors.New("State not matching")
	}

	rawIDToken, err := u.keycloakService.VerifyCallback(ctx, code)
	if err != nil {
		return nil, err
	}

	user, err := u.keycloakService.ConvertToken(ctx, rawIDToken)
	if err != nil {
		return nil, err
	}

	return user, nil
}