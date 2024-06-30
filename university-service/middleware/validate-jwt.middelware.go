package middleware

import (
	"context"
	"encoding/json"
	"fakultet-service/models"
	"fakultet-service/utils"
	"fmt"
	"log"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc, authServiceAddress string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := http.DefaultClient
		request, err := http.NewRequest("POST", fmt.Sprintf("%s/validate-jwt", authServiceAddress), nil)
		if err != nil {
			log.Fatalln(err.Error())
			utils.WriteErrorResp(err.Error(), 500, "/api/university", w)
			return
		}
		request.Header.Set("Authorization", r.Header.Get("Authorization"))
		response, err := client.Do(request)
		if err != nil {
			log.Fatalln(err.Error())
			utils.WriteErrorResp(err.Error(), 500, "/api/university", w)
			return
		}
		defer response.Body.Close()
		if response.StatusCode >= 400 {
			baseErrorResp := models.BaseErrorHttpResponse{}
			err := json.NewDecoder(response.Body).Decode(&baseErrorResp)
			if err != nil {
				log.Println(err)
				utils.WriteErrorResp(err.Error(), 500, "/api/university", w)
				return
			}
			utils.WriteErrorResp("Unauthorized", 401, "/api/university", w)
			return
		}
		baseHttpResponse := models.BaseHttpResponse{}
		err = json.NewDecoder(response.Body).Decode(&baseHttpResponse)
		if err != nil {
			utils.WriteErrorResp(err.Error(), 500, "/api/university", w)
			return
		}
		log.Println(baseHttpResponse)
		ctx := context.WithValue(r.Context(), "user", baseHttpResponse.Data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})

}
