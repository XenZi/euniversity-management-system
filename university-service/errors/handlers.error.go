package errors

import (
	"fakultet-service/models"

	"fmt"
	"go.mongodb.org/mongo-driver/mongo"

	"strings"
)

func HandleUniversityInsertError(err error, university models.University) (error, int) {
	if writeErr, ok := err.(mongo.WriteException); ok {
		for _, writeError := range writeErr.WriteErrors {
			if writeError.Code == 11000 {
				if strings.Contains(writeError.Message, "ID_1") {
					return fmt.Errorf("duplicate entity with universityID %s already exists", university.Id), 422
				}
			}
		}
	}
	return err, -1
}
