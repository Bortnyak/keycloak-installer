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
// Create realm according to the name specified in the config file (POST http://localhost:8080/auth/admin/realms)

// TODO:
// Turn on full scope (PUT http://localhost:8080/auth/admin/realms/NewTestRealm/clients/cad79dd6-ee36-45be-8ef5-e175c4ee0c41)
// {fullScopeAllowed: true}

// TODO:
// Update client (PUT http://localhost:8080/auth/admin/realms/NewTestRealm/clients/cad79dd6-ee36-45be-8ef5-e175c4ee0c41)
// {
// 	"id": "cad79dd6-ee36-45be-8ef5-e175c4ee0c41",
// 	"clientId": "admin-cli",
// 	"name": "${client_admin-cli}",
// 	"surrogateAuthRequired": false,
// 	"enabled": true,
// 	"alwaysDisplayInConsole": false,
// 	"clientAuthenticatorType": "client-secret",
// 	"redirectUris": [
// 	  "*"
// 	],
// 	"webOrigins": [],
// 	"notBefore": 0,
// 	"bearerOnly": false,
// 	"consentRequired": false,
// 	"standardFlowEnabled": true,
// 	"implicitFlowEnabled": false,
// 	"directAccessGrantsEnabled": true,
// 	"serviceAccountsEnabled": false,
// 	"publicClient": true,
// 	"frontchannelLogout": false,
// 	"protocol": "openid-connect",
// 	"attributes": {
// 	  "access.token.lifespan": 18000,
// 	  "client.session.idle.timeout": 36000,
// 	  "client.session.max.lifespan": 86400,
// 	  "saml.server.signature": "false",
// 	  "saml.server.signature.keyinfo.ext": "false",
// 	  "saml.assertion.signature": "false",
// 	  "saml.client.signature": "false",
// 	  "saml.encrypt": "false",
// 	  "saml.authnstatement": "false",
// 	  "saml.onetimeuse.condition": "false",
// 	  "saml_force_name_id_format": "false",
// 	  "saml.multivalued.roles": "false",
// 	  "saml.force.post.binding": "false",
// 	  "exclude.session.state.from.auth.response": "false",
// 	  "tls.client.certificate.bound.access.tokens": "false",
// 	  "display.on.consent.screen": "false"
// 	},
// 	"authenticationFlowBindingOverrides": {},
// 	"fullScopeAllowed": true,
// 	"nodeReRegistrationTimeout": 0,
// 	"defaultClientScopes": [
// 	  "web-origins",
// 	  "role_list",
// 	  "roles",
// 	  "profile",
// 	  "email"
// 	],
// 	"optionalClientScopes": [
// 	  "address",
// 	  "phone",
// 	  "offline_access",
// 	  "microprofile-jwt"
// 	],
// 	"access": {
// 	  "view": true,
// 	  "configure": true,
// 	  "manage": true
// 	},
// 	"authorizationServicesEnabled": ""
//   }
