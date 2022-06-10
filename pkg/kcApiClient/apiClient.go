package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	conf "github.com/Bortnyak/keycloak-installer/pkg/config"
	"github.com/Bortnyak/keycloak-installer/pkg/roles"
	"github.com/TwiN/go-color"
)

type KeycloakApiClient struct {
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

type KeycloakRealmClient struct {
	ClientId string `json:"clientId"`
	Name     string `json:"name"`
	Id       string `json:"id"`
}

const getTokenURLPath = "/realms/master/protocol/openid-connect/token"

type CreateRolePayload struct {
	Name string `json:"name"`
}

type CreateProtocolMapperPayload struct {
	Name           string `json:"name"`
	Protocol       string `json:"protocol"`
	ProtocolMapper string `json:"protocolMapper"`
}

func (apiClient *KeycloakApiClient) Authenticate() {
	config := conf.GetConf()
	reqForm := url.Values{
		"grant_type": {"password"},
		"client_id":  {config.Client},
		"username":   {config.AdminLogin},
		"password":   {config.AdminPassword},
	}

	reqURL := config.BaseURL + getTokenURLPath
	res, err := http.PostForm(reqURL, reqForm)
	if err != nil {
		fmt.Println("Failed to get auth token")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body")
	}

	fmt.Println(string(body))

	var keycloakResponse = new(KyecloakApiResponse)
	err1 := json.Unmarshal(body, &keycloakResponse)
	if err1 != nil {
		fmt.Println("Failed to unmarshal response body")
	}

	apiClient.token = keycloakResponse.Token
	println(color.Ize(color.Green, "Authenticated"))
}

func (apiClient *KeycloakApiClient) CreateRole(roleName string) {
	config := conf.GetConf()

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	roleNamePayload := &CreateRolePayload{Name: roleName}
	reqBody, err := json.Marshal(roleNamePayload)
	if err != nil {
		fmt.Println("Failed to marshal request body, e: ", err)
	}

	reqURL := config.BaseURL + "/admin/realms/" + config.Realm + "/roles"
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Failed to create new request: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiClient.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	fmt.Println(resp)
}

func (apiClinet *KeycloakApiClient) InitRoles() {
	var wg sync.WaitGroup
	for _, role := range roles.RoleIds {
		wg.Add(1)

		go func(role string) {
			defer wg.Done()
			apiClinet.CreateRole(role)
		}(role)
	}
	wg.Wait()
	println(color.Ize(color.Green, "Roles created"))
}

func (apiClient *KeycloakApiClient) CreateProtocolMapper() {
	realmClients := apiClient.GetClients()
	adminCliClient := getRightClientId(*realmClients)
	apiClient.createProtocolMapperForClient(adminCliClient)
	println(color.Ize(color.Green, "Protocol mapper created"))
}

func (apiClient *KeycloakApiClient) GetClients() *[]KeycloakRealmClient {
	config := conf.GetConf()

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	reqURL := config.BaseURL + "/admin/realms/" + config.Realm + "/clients"
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Println("Failed to create new request: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiClient.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Failed to get clinets: ", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to get response body: ", err)
	}

	var realmClients []KeycloakRealmClient
	err = json.Unmarshal(respBody, &realmClients)

	return &realmClients
}

func getRightClientId(clients []KeycloakRealmClient) string {
	conf := conf.GetConf()
	var id string

	for _, clientStruct := range clients {
		if clientStruct.ClientId == conf.Client {
			id = clientStruct.Id
		}
	}

	return id
}

func (apiClinet *KeycloakApiClient) createProtocolMapperForClient(clientId string) {
	config := conf.GetConf()

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	protocolPayload := &CreateProtocolMapperPayload{Name: "RealmAccessProtocolMapper", Protocol: "openid-connect", ProtocolMapper: "oidc-custom-roles-protocol-mapper"}
	reqBody, err := json.Marshal(protocolPayload)
	if err != nil {
		fmt.Println("Failed to marshal request body, e: ", err)
	}

	reqURL := config.BaseURL + "/admin/realms/" + config.Realm + "/clients/" + clientId + "/protocol-mappers/models"
	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Failed to create new request for protocol mapper: ", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiClinet.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	fmt.Println(resp)
}
