package errors

import (
	"auth/models"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func HandleInsertError(err error, user models.Citizen) (error, int) {
	if writeErr, ok := err.(mongo.WriteException); ok {
		for _, writeError := range writeErr.WriteErrors {
			if writeError.Code == 11000 {
				if strings.Contains(writeError.Message, "email_1") {
					return fmt.Errorf("Duplicate entity with email %s already exists", user.Email), 422
				} else if strings.Contains(writeError.Message, "pid_1") {
					return fmt.Errorf("Duplicate entity with username %s already exists", user.PersonalIdentificationNumber), 422
				}
			}
		}
	}
	return err, -1
}
