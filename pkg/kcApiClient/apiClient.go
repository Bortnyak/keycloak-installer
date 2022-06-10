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

const getTokenURLPath = "/realms/master/protocol/openid-connect/token"

type CreateRolePayload struct {
	Name string `json:"name"`
}

func (apiClinet *KeycloakApiClient) Authenticate() {
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

	apiClinet.token = keycloakResponse.Token
}

func (apiClinet *KeycloakApiClient) CreateRole(roleName string) error {
	fmt.Println("Inside roles")
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

	req.Header.Set("Authorization", "Bearer "+apiClinet.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	fmt.Println(resp)
	return nil
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
}
