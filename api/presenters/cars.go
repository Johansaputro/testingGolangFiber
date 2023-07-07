package presenters

import (
	"testingfiber/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CarName string             `json:"carName"`
	Company string             `json:"company"`
}

func CarSuccessResponse(data *entities.Car) *fiber.Map {
	car := Car{
		ID:      data.ID,
		CarName: data.CarName,
		Company: data.Company,
	}

	return &fiber.Map{
		"status": true,
		"data":   car,
		"error":  nil,
	}
}

func CarsSuccessResponse(datas *[]entities.Car) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   datas,
		"error":  nil,
	}
}

func CarErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err,
	}
}
