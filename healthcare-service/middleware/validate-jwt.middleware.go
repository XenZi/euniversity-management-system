package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"healthcare/models"
	"healthcare/utils"
	"log"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc, authServiceAdress string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := http.DefaultClient
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/validate-jwt", authServiceAdress), nil)
		if err != nil {
			utils.WriteErrorResp(err.Error(), 500, "/api/healthcare", w)
			return
		}
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		resp, err := client.Do(req)
		if err != nil {
			utils.WriteErrorResp(err.Error(), 500, "/api/healthcare", w)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			baseErrorResp := models.BaseErrorHttpResponse{}
			err := json.NewDecoder(resp.Body).Decode(&baseErrorResp)
			if err != nil {
				log.Println(err)
				utils.WriteErrorResp(err.Error(), 500, "/api/healthcare", w)
				return
			}
			utils.WriteErrorResp("Unathorized", 401, "/api/healthcare", w)
			return
		}
		baseHttpResponse := models.BaseHttpResponse{}
		err = json.NewDecoder(resp.Body).Decode(&baseHttpResponse)
		if err != nil {
			utils.WriteErrorResp(err.Error(), 500, "/api/healthcare", w)
			return
		}
		log.Println(baseHttpResponse)
		ctx := context.WithValue(r.Context(), "user", baseHttpResponse.Data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
