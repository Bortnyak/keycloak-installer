package main

import (
	apiclient "github.com/Bortnyak/keycloak-installer/pkg/kcApiClient"
)

func main() {
	keycloakAPI := &apiclient.KeycloakApiClient{}
	keycloakAPI.Authenticate()

	// all the necessary roles (ids) to map to roles from the database
	keycloakAPI.InitRoles()

	// protocol mapper to include custom a role in the access token
	keycloakAPI.CreateProtocolMapper()
}
