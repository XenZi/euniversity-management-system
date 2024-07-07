package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"healthcare/errors"
	"healthcare/models"
)

type DepartmentRepository struct {
	cli *mongo.Client
}

func NewDepartmentRepository(cli *mongo.Client) (*DepartmentRepository, error) {
	return &DepartmentRepository{
		cli: cli,
	}, nil
}

const (
	h2  = "department"
	dep = "department"
)

func (d DepartmentRepository) SaveDepartment(department models.Department) (*models.Department, *errors.ErrorStruct) {
	departmentCollection := d.cli.Database(h2).Collection(dep)
	insertedDepartment, err := departmentCollection.InsertOne(context.TODO(), department)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	department.ID = insertedDepartment.InsertedID.(primitive.ObjectID)
	return &department, nil
}

func (d DepartmentRepository) UpdateDepartment(department models.Department) (*models.Department, *errors.ErrorStruct) {
	departmentCollection := d.cli.Database(h2).Collection(dep)
	filter := bson.M{"_id": department.ID}
	update := bson.D{
		{"$set", bson.D{
			{"schedule", department.Schedule},
		}},
	}
	_, err := departmentCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return &department, nil
}

func (d DepartmentRepository) GetDepartmentByName(name string) (*models.Department, *errors.ErrorStruct) {
	departmentCollection := d.cli.Database(h2).Collection(dep)
	filter := bson.D{{"name", name}}
	var dept *models.Department
	err := departmentCollection.FindOne(context.TODO(), filter).Decode(&dept)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return dept, nil
}

func (d DepartmentRepository) GetAllDepartments() ([]*models.Department, *errors.ErrorStruct) {
	departmentCollection := d.cli.Database(h2).Collection(dep)
	filter := bson.M{}
	var depts []*models.Department
	cursor, err := departmentCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var dept *models.Department
		if err := cursor.Decode(&dept); err != nil {
			return nil, errors.NewError(err.Error(), 500)
		}
		depts = append(depts, dept)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.NewError(err.Error(), 500)
	}
	return depts, nil
}
