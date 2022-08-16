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

// TODO:
// Turn on full scope (PUT http://localhost:8080/auth/admin/realms/NewTestRealm/clients/cad79dd6-ee36-45be-8ef5-e175c4ee0c41)
// {fullScopeAllowed: true}
