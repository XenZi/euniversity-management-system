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

// RECORDS

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

func (h HealthcareRepository) GetAllRecords() ([]*models.Record, *errors.ErrorStruct) {
	recordCollection := h.cli.Database(h1).Collection(rec)
	var recs []*models.Record
	cursor, err := recordCollection.Find(context.TODO(), nil)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var reco models.Record
		if err := cursor.Decode(&reco); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		recs = append(recs, &reco)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return recs, nil
}

// CERTIFICATES

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

// REFERRALS

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

// REPORTS

func (h HealthcareRepository) SaveReport(report models.Report) (*models.Report, *errors.ErrorStruct) {
	reportCollection := h.cli.Database(h1).Collection(rep)
	insertedReport, err := reportCollection.InsertOne(context.TODO(), report)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	report.ID = insertedReport.InsertedID.(primitive.ObjectID)
	return &report, nil
}

func (h HealthcareRepository) GetReportByID(id string) (*models.Report, *errors.ErrorStruct) {
	reportCollection := h.cli.Database(h1).Collection(rep)
	foundId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{"_id", foundId}}
	var report *models.Report
	erro := reportCollection.FindOne(context.TODO(), filter).Decode(&report)
	if erro != nil {
		err, status := errors.HandleNoDocumentsError(erro, id)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	return report, nil
}

// APPOINTMENTS

func (h HealthcareRepository) SaveAppointment(appointment models.Appointment) (*models.Appointment, *errors.ErrorStruct) {
	appointmentCollection := h.cli.Database(h1).Collection(app)
	insertedAppointment, err := appointmentCollection.InsertOne(context.TODO(), appointment)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	appointment.ID = insertedAppointment.InsertedID.(primitive.ObjectID)
	erro := h.updateRecordSideEffect(appointment.PatientID, nil, nil, nil, &appointment)
	if erro != nil {
		return nil, erro
	}
	return &appointment, nil
}

func (h HealthcareRepository) UpdateAppointment(appointment models.Appointment) (*models.Appointment, *errors.ErrorStruct) {
	appointmentCollection := h.cli.Database(h1).Collection(app)
	filter := bson.M{"_id": appointment.ID}
	update := bson.D{
		{"$set", bson.D{
			{"appointmentStatus", appointment.AppointmentStatus},
			{"report", appointment.Report},
		}},
	}
	_, err := appointmentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	erro := h.updateByDepth(appointment.PatientID, nil, &appointment)
	if erro != nil {
		return nil, erro
	}
	return &appointment, nil
}

func (h HealthcareRepository) GetAllAppointmentsByDoctorID(doctorID string) ([]*models.Appointment, *errors.ErrorStruct) {
	appointmentCollection := h.cli.Database(h1).Collection(app)
	filter := bson.M{"doctorID": doctorID}
	var appointments []*models.Appointment
	cursor, err := appointmentCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var appo models.Appointment
		if err := cursor.Decode(&appo); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		appointments = append(appointments, &appo)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return appointments, nil
}

func (h HealthcareRepository) GetAppointmentByID(id string) (*models.Appointment, *errors.ErrorStruct) {
	appointmentCollection := h.cli.Database(h1).Collection(app)
	foundId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{"_id", foundId}}
	var appointment *models.Appointment
	erro := appointmentCollection.FindOne(context.TODO(), filter).Decode(&appointment)
	if erro != nil {
		err, status := errors.HandleNoDocumentsError(erro, id)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	return appointment, nil
}

// PRESCRIPTIONS

func (h HealthcareRepository) SavePrescription(prescription models.Prescription) (*models.Prescription, *errors.ErrorStruct) {
	prescriptionCollection := h.cli.Database(h1).Collection(pre)
	insertedPrescription, err := prescriptionCollection.InsertOne(context.TODO(), prescription)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	prescription.ID = insertedPrescription.InsertedID.(primitive.ObjectID)
	erro := h.updateRecordSideEffect(prescription.PatientID, nil, nil, &prescription, nil)
	if erro != nil {
		return nil, erro
	}
	return &prescription, nil
}

func (h HealthcareRepository) GetPrescriptionByID(id string) (*models.Prescription, *errors.ErrorStruct) {
	prescriptionCollection := h.cli.Database(h1).Collection(pre)
	foundId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	filter := bson.D{{"_id", foundId}}
	var prescription *models.Prescription
	erro := prescriptionCollection.FindOne(context.TODO(), filter).Decode(&prescription)
	if erro != nil {
		err, status := errors.HandleNoDocumentsError(erro, id)
		if status == -1 {
			status = 500
		}
		return nil, errors.NewError(err.Error(), status)
	}
	return prescription, nil
}

func (h HealthcareRepository) UpdatePrescription(prescription models.Prescription) (*models.Prescription, *errors.ErrorStruct) {
	prescriptionCollection := h.cli.Database(h1).Collection(pre)
	filter := bson.M{"_id": prescription.ID}
	update := bson.D{
		{"$set", bson.D{
			{"prescriptionStatus", prescription.Status},
		}},
	}
	_, err := prescriptionCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	erro := h.updateByDepth(prescription.PatientID, &prescription, nil)
	if erro != nil {
		return nil, erro
	}
	return &prescription, nil
}

// LOGIC UTILS

func (h HealthcareRepository) updateByDepth(id string, pres *models.Prescription, appo *models.Appointment) *errors.ErrorStruct {
	record, err := h.GetRecordByPatientID(id)
	if err != nil {
		return err
	}
	counter := 0
	if pres != nil {
		for _, prescription := range record.Prescriptions {
			if prescription.ID == pres.ID {
				prescription = *pres
			}
		}
		counter += 1
	}
	if appo != nil {
		for _, appointment := range record.Appointments {
			if appointment.ID == appo.ID {
				appointment = *appo
			}
		}
		counter += 1
	}
	if counter > 0 {
		_, err = h.UpdateRecord(*record)
	}
	return nil

}

func (h HealthcareRepository) updateRecordSideEffect(id string, cert *models.Certificate, ref *models.Referral, pres *models.Prescription, appo *models.Appointment) *errors.ErrorStruct {
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
	if appo != nil {
		_ = append(record.Appointments, *appo)
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
