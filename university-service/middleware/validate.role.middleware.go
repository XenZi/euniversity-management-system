package middleware

import (
	"fakultet-service/models"
	"fakultet-service/utils"
	"fmt"
	"net/http"
)

func ValidateRole(next http.HandlerFunc, requiredRole string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contextValue, ok := r.Context().Value("user").(map[string]interface{})
		if !ok {
			fmt.Println("Map not found in context or wrong type")
			return
		}

		rawRoles, ok := contextValue["roles"].([]interface{})
		if !ok {
			fmt.Println("Roles not found or wrong type")
			return
		}

		roles := make([]string, len(rawRoles))
		for i, v := range rawRoles {
			roles[i], ok = v.(string)
			if !ok {
				fmt.Println("Role value is not a string")
				return
			}
		}
		person := models.Citizen{
			PersonalIdentificationNumber: contextValue["personalIdentificationNumber"].(string),
			FullName:                     contextValue["fullName"].(string),
			Roles:                        roles,
		}
		if person.Roles[0] != requiredRole {
			utils.WriteErrorResp("Unauthorized", 401, "/api/university", w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
