package main

import apiclient "github.com/Bortnyak/keycloak-installer/pkg/kcApiClient"

func main() {
	c := &apiclient.KeycloakApiClient{}
	c.CreateRole("name")
}
