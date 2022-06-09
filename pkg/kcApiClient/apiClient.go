package apiclient

type KeycloakApiClient struct {
	baseURL             string
	masterRealmURL      string
	masterRealmTokenURL string
	realmName           string
	realmURL            string
	roleEndpoint        string
	token               string
}

type KyecloakApiResponse struct {
	Token                 string `json:"access_token"`
	Expires               int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	TokenType             string `json:"token_type"`
	NotBeforePolicy       int    `json:"not-before-policy"`
	SessionState          string `json:"sessoin_state"`
	Scope                 string `json:"scope"`
}

type CreateRolePayload struct {
	Name string `json:"name"`
}
