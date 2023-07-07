package cars

import "testingfiber/pkg/entities"

type Service interface {
	InsertCarService(car *entities.Car) (*entities.Car, error)
	CheckCarService() (*[]entities.Car, error)
	UpdateCarService(car *entities.Car) (*entities.Car, error)
	RemoveCarService(ID string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertCarService(car *entities.Car) (*entities.Car, error) {
	return s.repository.InsertCar(car)
}

func (s *service) CheckCarService() (*[]entities.Car, error) {
	return s.repository.CheckCar()
}

func (s *service) UpdateCarService(car *entities.Car) (*entities.Car, error) {
	return s.repository.UpdateCar(car)
}

func (s *service) RemoveCarService(ID string) error {
	return s.repository.DeleteCar(ID)
}
