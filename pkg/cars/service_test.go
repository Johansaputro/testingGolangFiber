package cars

import (
	"testing"
	"testingfiber/pkg/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) InsertCar(car *entities.Car) (*entities.Car, error) {
	args := m.Called(car)
	result := args.Get(0)
	err := args.Error(1)

	if result != nil {
		return result.(*entities.Car), err
	}

	return nil, err
}

func (m *mockRepository) CheckCar() (*[]entities.Car, error) {
	args := m.Called()
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*[]entities.Car), err
	}
	return nil, err
}

func (m *mockRepository) UpdateCar(car *entities.Car) (*entities.Car, error) {
	args := m.Called(car)
	result := args.Get(0)
	err := args.Error(1)
	if result != nil {
		return result.(*entities.Car), err
	}
	return nil, err
}

func (m *mockRepository) DeleteCar(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func TestInsertCarService(t *testing.T) {
	repo := new(mockRepository)
	service := NewService(repo)

	car := &entities.Car{
		CarName: "Mazda",
	}

	expectedCar := &entities.Car{
		CarName: "Mazda",
	}

	repo.On("InsertCar", car).Return(expectedCar, nil)

	result, err := service.InsertCarService(car)

	assert.NoError(t, err)
	assert.Equal(t, expectedCar, result)

	repo.AssertExpectations(t)
}

func TestCheckCarService(t *testing.T) {
	repo := new(mockRepository)
	service := NewService(repo)

	cars := &[]entities.Car{
		{
			CarName: "Mazda",
		},
		{
			CarName: "Toyota",
		},
	}

	expectedCar := &cars

	repo.On("CheckCar").Return(expectedCar, nil)

	result, err := service.CheckCarService()

	assert.NoError(t, err)
	assert.Equal(t, expectedCar, result)

	repo.AssertExpectations(t)
}

func TestUpdateCarService(t *testing.T) {
	repo := new(mockRepository)
	service := NewService(repo)

	car := &entities.Car{
		CarName: "Mazda",
	}

	expectedCar := &entities.Car{
		CarName: "Mazda",
	}

	repo.On("UpdateCar", car).Return(expectedCar, nil)

	result, err := service.UpdateCarService(car)

	assert.NoError(t, err)
	assert.Equal(t, expectedCar, result)

	repo.AssertExpectations(t)
}

func TestRemoveCarService(t *testing.T) {
	repo := new(mockRepository)
	service := NewService(repo)

	ID := "123"

	repo.On("DeleteCar", ID).Return(nil)

	err := service.RemoveCarService(ID)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}
