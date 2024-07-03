package clients

import (
	"encoding/json"
	"fmt"
	"healthcare/errors"
	"healthcare/models"
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

func (uc UniversityClient) CheckIfStudent(id string) (bool, *errors.ErrorStruct) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/student/%s", uc.address, id), nil)
	if err != nil {
		log.Fatalln(err.Error())
		return false, errors.NewError(err.Error(), 500)
	}
	resp, erro := uc.client.Do(req)
	if erro != nil {
		log.Fatalln(erro.Error())
		return false, errors.NewError(erro.Error(), 500)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true, nil
	}
	baseResp := models.BaseErrorHttpResponse{}
	err1 := json.NewDecoder(resp.Body).Decode(&baseResp)
	if err1 != nil {
		return false, errors.NewError(err1.Error(), 500)
	}
	return false, errors.NewError(baseResp.Error, baseResp.Status)
}
