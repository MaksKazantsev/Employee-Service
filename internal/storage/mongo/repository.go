package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/MaksKazantsev/mongodb/internal/models"
	"github.com/MaksKazantsev/mongodb/internal/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRepository(cl *mongo.Client) storage.Storage {
	return &repository{
		db: cl.Database("employeeDB"),
	}
}

type repository struct {
	db *mongo.Database
}

func (r repository) Get(ctx context.Context, id int) (models.Employee, error) {
	var employee models.Employee
	coll := r.db.Collection("employee")
	filter := bson.M{
		"_id": id,
	}
	err := coll.FindOne(ctx, filter).Decode(&employee)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return employee, fmt.Errorf("no documents!, error: %v", err)
		}
		return employee, fmt.Errorf("failed to exectute get query, error: %v", err)
	}
	return employee, nil
}

func (r repository) Add(ctx context.Context, e *models.Employee) error {
	coll := r.db.Collection("employee")

	_, err := coll.InsertOne(ctx, e)
	if err != nil {
		return fmt.Errorf("failed to exectute insert query, error: %v", err)
	}

	return nil
}

func (r repository) Delete(ctx context.Context, id int) error {
	coll := r.db.Collection("employee")

	filter := bson.M{
		"_id": id,
	}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to exectute delete query, error: %v", err)
	}

	return nil
}

func (r repository) Update(ctx context.Context, id int, e models.Employee) error {

	coll := r.db.Collection("employee")

	update := bson.D{
		{Key: "$set", Value: e},
	}

	filter := bson.D{
		{Key: "_id", Value: id},
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to exectute update query, error: %v", err)
	}
	return nil
}

func (r repository) GetAll(ctx context.Context) ([]models.Employee, error) {

	var employees []models.Employee
	coll := r.db.Collection("employee")

	filter := bson.D{}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to exectute get query, error: %v", err)
	}

	err = cur.All(ctx, &employees)
	if err != nil {
		return nil, fmt.Errorf("failed to decode, error: %v", err)
	}

	return employees, nil
}

func (r repository) DeleteAll(ctx context.Context) (error, int64) {

	coll := r.db.Collection("employee")

	filter := bson.D{}

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to exectute delete query, error: %v", err), 0
	}

	return nil, res.DeletedCount
}
