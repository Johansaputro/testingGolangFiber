package cars

import (
	"context"
	"testing"
	"testingfiber/pkg/entities"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

type RepositoryWrapper interface {
	InsertCar(car *entities.Car) (*entities.Car, error)
	CheckCar() (*[]entities.Car, error)
	UpdateCar(book *entities.Car) (*entities.Car, error)
	DeleteCar(ID string) error
}

type repositoryWrapper struct {
	Collection *MockCollection
}

func NewRepoWrapper(collection *MockCollection) RepositoryWrapper {
	return &repositoryWrapper{
		Collection: collection,
	}
}

func (r *repositoryWrapper) InsertCar(car *entities.Car) (*entities.Car, error) {
	car.ID = primitive.NewObjectID()
	car.MadeAt = time.Now()
	car.SoldAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), car)

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repositoryWrapper) CheckCar() (*[]entities.Car, error) {
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

func (r *repositoryWrapper) UpdateCar(car *entities.Car) (*entities.Car, error) {
	car.SoldAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": car.ID}, bson.M{"$set": car})

	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repositoryWrapper) DeleteCar(ID string) error {
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

func TestInsertCar(t *testing.T) {
	car := &entities.Car{}
	mockCollection := &MockCollection{}

	// Set up the expected behavior of the mock collection
	mockCollection.On("InsertOne", mock.Anything, car).Return(&mongo.InsertOneResult{}, nil)

	repo := NewRepoWrapper(mockCollection)

	insertedCar, err := repo.InsertCar(car)

	assert.NoError(t, err)
	assert.NotNil(t, insertedCar)
	mockCollection.AssertExpectations(t)
}

func TestCheckCar(t *testing.T) {
	// Create a mock cursor with sample car documents
	mockCursor := &MockCursor{}
	mockCursor.On("Next", mock.Anything).Return(true)
	mockCursor.On("Decode", mock.AnythingOfType("*entities.Car")).Return(nil)
	mockCursor.On("Next", mock.Anything).Return(false)
	mockCursor.On("Err").Return(nil)
	mockCursor.On("Close", mock.Anything).Return(nil)

	mockCollection := &MockCollection{}
	mockCollection.On("Find", mock.Anything, bson.D{}).Return(mockCursor, nil)

	repo := NewRepoWrapper(mockCollection)

	cars, err := repo.CheckCar()

	assert.NoError(t, err)
	assert.NotNil(t, cars)
	mockCollection.AssertExpectations(t)
	mockCursor.AssertExpectations(t)
}

func TestUpdateCar(t *testing.T) {
	car := &entities.Car{ID: primitive.NewObjectID()}
	mockCollection := &MockCollection{}

	// Set up the expected behavior of the mock collection
	mockCollection.On("UpdateOne", mock.Anything, bson.M{"_id": car.ID}, bson.M{"$set": car}).Return(&mongo.UpdateResult{}, nil)

	repo := NewRepoWrapper(mockCollection)

	updatedCar, err := repo.UpdateCar(car)

	assert.NoError(t, err)
	assert.NotNil(t, updatedCar)
	mockCollection.AssertExpectations(t)
}

func TestDeleteCar(t *testing.T) {
	carID := primitive.NewObjectID().Hex()
	mockCollection := &MockCollection{}

	// Set up the expected behavior of the mock collection
	mockCollection.On("DeleteOne", mock.Anything, bson.M{"_id": carID}).Return(&mongo.DeleteResult{}, nil)

	repo := NewRepoWrapper(mockCollection)

	err := repo.DeleteCar(carID)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}
