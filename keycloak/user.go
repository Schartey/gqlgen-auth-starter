package keycloak

type User struct {
	UserToken	UserToken
	RawIDToken  interface{}
	UserRoles	[]*UserRole
}

type UserToken struct {
	Jti                string
	Exp                int
	Nbf                int
	Iat                int
	Iss                string
	Aud                string
	Sub                string
	Typ                string
	Azp                string
	Auth_time          int
	Session_state      string
	Acr                string
	Email_verified     bool
	Name               string
	Preferred_username string
	Given_name         string
	Family_name        string
	Email              string

}

type UserRole struct {
	Name string
}