package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthServiceClient struct {
	address string
	client  *http.Client
}

func NewAuthServiceClient(address string, client *http.Client) *AuthServiceClient {
	return &AuthServiceClient{
		address: address,
		client:  client,
	}
}

func (as *AuthServiceClient) AddRoles(pin string, roles []string) error {

	type AddingRoles struct {
		PersonalIdentificationNumber string   `json:"pin"`
		Roles                        []string `json:"roles"`
	}
	payload := AddingRoles{
		PersonalIdentificationNumber: pin,
		Roles:                        roles,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/addRoles", as.address)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := as.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
