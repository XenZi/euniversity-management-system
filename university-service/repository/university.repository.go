package repository

import (
	"context"
	"fakultet-service/errors"
	"fakultet-service/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UniversityRepository struct {
	cli *mongo.Client
}

func NewUniversityRepository(cli *mongo.Client) (*UniversityRepository, error) {
	return &UniversityRepository{
		cli: cli,
	}, nil
}

func (u UniversityRepository) SaveUniversity(university models.University) (*models.University, *errors.ErrorStruct) {
	universityCollection := u.cli.Database("university").Collection("university_collection")
	insertResult, err := universityCollection.InsertOne(context.TODO(), university)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	university.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &university, nil

}
func (u UniversityRepository) SaveProfessor(professor models.Professor) (*models.Professor, *errors.ErrorStruct) {
	professorCollection := u.cli.Database("university").Collection("professor")
	insertResult, err := professorCollection.InsertOne(context.TODO(), professor)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	professor.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &professor, nil
}

func (u UniversityRepository) SaveStudent(student models.Student) (*models.Student, *errors.ErrorStruct) {
	studentCollection := u.cli.Database("university").Collection("student")
	insertResult, err := studentCollection.InsertOne(context.TODO(), student)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	student.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &student, nil
}

func (u UniversityRepository) SaveScholarship(scholarship models.Scholarship) (*models.Scholarship, *errors.ErrorStruct) {
	scholarshipCollection := u.cli.Database("university").Collection("scholarship")
	insertResult, err := scholarshipCollection.InsertOne(context.TODO(), scholarship)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	scholarship.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &scholarship, nil
}

func (u UniversityRepository) SaveStateExamApplication(application models.StateExamApplication) (*models.StateExamApplication, *errors.ErrorStruct) {
	applicationCollection := u.cli.Database("university").Collection("state_exam_application")
	insertResult, err := applicationCollection.InsertOne(context.TODO(), application)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	application.ID = insertResult.InsertedID.(primitive.ObjectID)
	return &application, nil
}

func (u UniversityRepository) FindStudentById(personalIdentificationNumber string) (*models.Student, *errors.ErrorStruct) {
	studentCollection := u.cli.Database("university").Collection("student")
	var student models.Student
	err := studentCollection.FindOne(context.TODO(), bson.M{"personalIdentificationNumber": personalIdentificationNumber}).Decode(&student)
	if err != nil {
		return nil, errors.NewError(err.Error(), 400)
	}
	fmt.Println(student)
	return &student, nil
}

func (u UniversityRepository) FindProfessor(personalIdentificationNumber string) (*models.Professor, *errors.ErrorStruct) {
	professorCollection := u.cli.Database("university").Collection("professor")
	var professor models.Professor
	err := professorCollection.FindOne(context.TODO(), bson.M{"personalIdentificationNumber": personalIdentificationNumber}).Decode(&professor)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	fmt.Println(professor)
	return &professor, nil

}

func (u UniversityRepository) FindScholarship(id primitive.ObjectID) (*models.Scholarship, *errors.ErrorStruct) {
	scholarshipCollection := u.cli.Database("university").Collection("scholarship")
	var scholarship models.Scholarship
	err := scholarshipCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&scholarship)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	fmt.Println(scholarship)
	return &scholarship, nil
}

func (u UniversityRepository) UpdateStudent(student models.Student) (*models.Student, *errors.ErrorStruct) {
	studentCollection := u.cli.Database("university").Collection("student")
	filter := bson.M{"personalIdentificationNumber": student.PersonalIdentificationNumber}
	update := bson.M{"$set": student}

	updateResult, err := studentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	if updateResult.ModifiedCount == 0 {
		return nil, errors.NewError("No student was updated", 400)
	}

	return &student, nil
}

func (u UniversityRepository) UpdateProfessor(professor models.Professor) (*models.Professor, *errors.ErrorStruct) {

	professorCollection := u.cli.Database("university").Collection("professor")
	filter := bson.M{"personalIdentificationNumber": professor.PersonalIdentificationNumber}
	update := bson.M{"$set": professor}

	updateResult, err := professorCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	if updateResult.ModifiedCount == 0 {
		return nil, errors.NewError("No professor was updated", 400)
	}
	return &professor, nil
}

func (u UniversityRepository) DeleteStudent(personalIdentificationNumber string) (*models.Student, *errors.ErrorStruct) {
	student, err := u.FindStudentById(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	studentCollection := u.cli.Database("university").Collection("student")
	filter := bson.M{"personalIdentificationNumber": personalIdentificationNumber}
	_, errFromDelete := studentCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return student, nil
}

func (u UniversityRepository) DeleteProfessor(personalIdentificationNumber string) (*models.Professor, *errors.ErrorStruct) {
	professor, err := u.FindProfessor(personalIdentificationNumber)
	if err != nil {
		return nil, err
	}
	professorCollection := u.cli.Database("university").Collection("professor")
	filter := bson.M{"personalIdentificationNumber": personalIdentificationNumber}
	_, errFromDelete := professorCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return professor, nil
}

func (u UniversityRepository) DeleteScholarship(id primitive.ObjectID) (*models.Scholarship, *errors.ErrorStruct) {
	scholarship, err := u.FindScholarship(id)
	if err != nil {
		return nil, err
	}
	scholarshipCollection := u.cli.Database("university").Collection("scholarship")
	filter := bson.M{"_id": id}
	_, errFromDelete := scholarshipCollection.DeleteOne(context.TODO(), filter)
	if errFromDelete != nil {
		return nil, errors.NewError(errFromDelete.Error(), 500)
	}
	return scholarship, nil
}
