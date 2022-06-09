package roles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	apiclient "github.com/Bortnyak/keycloak-installer/pkg/kcApiClient"
)

var roleIds = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func (c *apiclient.KeycloakApiClient) CreateRole(roleName string) error {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	roleNamePayload := &apiclient.CreateRolePayload{Name: roleName}
	reqBody, err := json.Marshal(roleNamePayload)
	if err != nil {
		fmt.Println("Failed to marshal request body, e: ", err)
	}

	//
	reqURL := base
}
