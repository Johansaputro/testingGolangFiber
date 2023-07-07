package cars

import (
	"context"
	"testingfiber/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertCar(car *entities.Car) (*entities.Car, error)
	CheckCar() (*[]entities.Car, error)
	UpdateCar(book *entities.Car) (*entities.Car, error)
	DeleteCar(ID string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) InsertCar(car *entities.Car) (*entities.Car, error) {
	car.ID = primitive.NewObjectID()
	car.MadeAt = time.Now()
	car.SoldAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), car)

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repository) CheckCar() (*[]entities.Car, error) {
	var cars []entities.Car
	cursor, err := r.Collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var car entities.Car
		_ = cursor.Decode(&car)

		cars = append(cars, car)
	}

	return &cars, nil
}

func (r *repository) UpdateCar(car *entities.Car) (*entities.Car, error) {
	car.SoldAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": car.ID}, bson.M{"$set": car})

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repository) DeleteCar(ID string) error {
	carId, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": carId})

	if err != nil {
		return err
	}

	return nil
}
