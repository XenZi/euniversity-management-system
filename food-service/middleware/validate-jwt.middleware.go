package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"food/models"
	"food/utils"
	"log"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc, foodServiceAdress string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("PRVO")
		client := http.DefaultClient
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/validate-jwt", foodServiceAdress), nil)
		if err != nil {
			log.Fatalln(err.Error())
			utils.WriteErrorResp(err.Error(), 500, "/api/food", w)
			return
		}
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln("TTST", err.Error())
			utils.WriteErrorResp(err.Error(), 500, "/api/dorm", w)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 400 {
			baseErrorResp := models.BaseErrorHttpResponse{}
			err := json.NewDecoder(resp.Body).Decode(&baseErrorResp)
			if err != nil {
				log.Println(err)
				utils.WriteErrorResp(err.Error(), 500, "/api/dorm", w)
				return
			}
			utils.WriteErrorResp("Unathorized", 401, "/api/dorm", w)
			return
		}
		baseHttpResponse := models.BaseHttpResponse{}
		err = json.NewDecoder(resp.Body).Decode(&baseHttpResponse)
		if err != nil {
			utils.WriteErrorResp(err.Error(), 500, "/api/dorm", w)
			return
		}
		log.Println(baseHttpResponse)
		ctx := context.WithValue(r.Context(), "user", baseHttpResponse.Data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
