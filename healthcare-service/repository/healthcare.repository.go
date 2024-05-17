package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare/errors"
	"healthcare/models"
)

type HealthcareRepository struct {
	cli *mongo.Client
}

func NewHealthcareRepository(cli *mongo.Client) (*HealthcareRepository, error) {
	return &HealthcareRepository{
		cli: cli,
	}, nil
}

const (
	h1   = "healthcare"
	rec  = "records"
	cert = "certificate"
	pre  = "prescription"
	ref  = "referral"
	app  = "appointment"
	rep  = "report"
)

func (h HealthcareRepository) SaveRecord(record models.Record) (*models.Record, *errors.ErrorStruct) {
	recordCollection := h.cli.Database(h1).Collection(rec)
	insertedRecord, err := recordCollection.InsertOne(context.TODO(), record)
	if err != nil {
		err, status := errors.HandleRecordInsertError(err, record)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	record.ID = insertedRecord.InsertedID.(primitive.ObjectID)
	return &record, nil
}

func (h HealthcareRepository) SaveReferral(referral models.Referral) (*models.Referral, *errors.ErrorStruct) {
	referralCollection := h.cli.Database(h1).Collection(ref)
	insertedReferral, err := referralCollection.InsertOne(context.TODO(), referral)
	if err != nil {
		err, status := errors.HandleReferralInsertError(err, referral)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	referral.ID = insertedReferral.InsertedID.(primitive.ObjectID)
	erro := h.updateRecordSideEffect(referral.PatientID, nil, &referral, nil, nil)
	if erro != nil {
		return nil, erro
	}
	return &referral, nil
}

func (h HealthcareRepository) SaveCertificate(certificate models.Certificate) (*models.Certificate, *errors.ErrorStruct) {
	certificateCollection := h.cli.Database(h1).Collection(cert)
	insertedCertificate, err := certificateCollection.InsertOne(context.TODO(), certificate)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	certificate.ID = insertedCertificate.InsertedID.(primitive.ObjectID)
	erro := h.updateRecordSideEffect(certificate.PatientID, &certificate, nil, nil, nil)
	if erro != nil {
		return nil, erro
	}
	return &certificate, nil
}

func (h HealthcareRepository) SaveReport(report models.Report) (*models.Report, *errors.ErrorStruct) {
	reportCollection := h.cli.Database(h1).Collection(rep)
	insertedReport, err := reportCollection.InsertOne(context.TODO(), report)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	report.ID = insertedReport.InsertedID.(primitive.ObjectID)
	return &report, nil
}

func (h HealthcareRepository) UpdateRecord(record models.Record) (*models.Record, *errors.ErrorStruct) {
	recordCollection := h.cli.Database(h1).Collection(rec)
	filter := bson.M{"_id": record.ID}
	update := bson.D{
		{"$set", bson.D{
			{"certificate", record.Certificate},
			{"prescriptions", record.Prescriptions},
			{"referrals", record.Referrals},
			{"appointments", record.Appointments},
		}},
	}
	_, err := recordCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return &record, nil
}

func (h HealthcareRepository) GetRecordByPatientID(patientID string) (*models.Record, *errors.ErrorStruct) {
	recordCollection := h.cli.Database(h1).Collection(rec)

	filter := bson.D{{"patientID", patientID}}

	var record *models.Record
	err := recordCollection.FindOne(context.TODO(), filter).Decode(&record)

	if err != nil {
		fmt.Println(err.Error())

		err, status := errors.HandleNoDocumentsError(err, patientID)
		if status == -1 {
			status = 500
			fmt.Println(err.Error())
		}
		if status == 404 {
			newRec := models.Record{
				PatientID: patientID,
			}
			inserted, erro := h.SaveRecord(newRec)
			if erro != nil {
				return nil, erro
			}
			return inserted, nil
		}
		return nil, errors.NewError(err.Error(), status)
	}
	return record, nil
}

func (h HealthcareRepository) GetReferralByID(id string) (*models.Referral, *errors.ErrorStruct) {
	referralCollection := h.cli.Database(h1).Collection(ref)
	foundId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{"_id", foundId}}
	var referral *models.Referral
	erro := referralCollection.FindOne(context.TODO(), filter).Decode(&referral)
	if erro != nil {
		err, status := errors.HandleNoDocumentsError(erro, id)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	return referral, nil
}

func (h HealthcareRepository) updateRecordSideEffect(id string, cert *models.Certificate, ref *models.Referral, pres *models.Prescription, app *models.Appointment) *errors.ErrorStruct {
	record, err := h.GetRecordByPatientID(id)
	counter := 0
	if err != nil {
		return err
	}
	if ref != nil {
		_ = append(record.Referrals, *ref)
		counter += 1
	}
	if pres != nil {
		_ = append(record.Prescriptions, *pres)
		counter += 1
	}
	if app != nil {
		_ = append(record.Appointments, *app)
		counter += 1
	}
	if cert != nil {
		record.Certificate = *cert
		counter += 1
	}
	if counter > 0 {
		_, err = h.UpdateRecord(*record)
	}
	return nil
}
