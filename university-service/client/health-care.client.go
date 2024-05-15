package client

import (
	"encoding/json"
	"fakultet-service/errors"
	"fakultet-service/models"
	"fmt"
	"log"
	"net/http"
)

type HealthCareClient struct {
	address string
	client  *http.Client
}

func NewHealthCareClient(host, port string, client *http.Client) *HealthCareClient {
	return &HealthCareClient{
		address: fmt.Sprintf("http://%s:%s", host, port),
		client:  client,
	}
}
func (hc HealthCareClient) GetUserHealthStatusConfirmation(userID string) (bool, *errors.ErrorStruct) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/getCertificate/%s", hc.address, userID), nil)
	if err != nil {
		log.Fatalln(err.Error())
		return false, errors.NewError(err.Error(), 500)
	}
	resp, err := hc.client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
		return false, errors.NewError(err.Error(), 500)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		baseErrorResp := models.BaseErrorHttpResponse{}
		err := json.NewDecoder(resp.Body).Decode(&baseErrorResp)
		if err != nil {
			return false, errors.NewError(err.Error(), 500)
		}
		log.Println(baseErrorResp)
		return false, nil
	}
	baseHttpResponse := models.BaseHttpResponse{}
	err = json.NewDecoder(resp.Body).Decode(&baseHttpResponse)
	if err != nil {
		return false, errors.NewError(err.Error(), 500)
	}
	log.Println(baseHttpResponse)
	return true, nil
}
