package client

import (
	"dorm-service/errors"
	"dorm-service/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UniversityClient struct {
	address string
	client  *http.Client
}

func NewUniversityClient(host, port string, client *http.Client) *UniversityClient {
	return &UniversityClient{
		address: fmt.Sprintf("http://%s:%s", host, port),
		client:  client,
	}
}

func (uc UniversityClient) VerifyUserIntegrityWithUniversity(userID string) (bool, *errors.ErrorStruct) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/getCertificate/%s", uc.address, userID), nil)
	if err != nil {
		log.Fatalln(err.Error())
		return false, errors.NewError(err.Error(), 500)
	}
	resp, err := uc.client.Do(req)
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
