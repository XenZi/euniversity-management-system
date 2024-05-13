package errors

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare/models"
	"strings"
)

func HandleRecordInsertError(err error, record models.Record) (error, int) {
	if writeErr, ok := err.(mongo.WriteException); ok {
		for _, writeError := range writeErr.WriteErrors {
			if writeError.Code == 11000 {
				if strings.Contains(writeError.Message, "patientID_1") {
					return fmt.Errorf("duplicate entity with patientID %s already exists", record.PatientID), 422
				}
			}
		}
	}
	return err, -1
}

func HandleReferralInsertError(err error, referral models.Referral) (error, int) {
	if writeErr, ok := err.(mongo.WriteException); ok {
		for _, writeError := range writeErr.WriteErrors {
			if writeError.Code == 11000 {
				if strings.Contains(writeError.Message, "patientID_1") {
					return fmt.Errorf("duplicate entity with patientID %s already exists", referral.PatientID), 422
				}
			}
		}
	}
	return err, -1
}

func HandleNoDocumentsError(err error, id string) (error, int) {
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("No entity with id %s found", id), 404
	}
	return err, -1
}
