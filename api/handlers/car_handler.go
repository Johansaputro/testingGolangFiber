package handlers

import (
	"errors"
	"net/http"
	"testingfiber/api/presenters"
	"testingfiber/pkg/cars"
	"testingfiber/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

func AddCar(service cars.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Car
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.CarErrorResponse(err))
		}

		if requestBody.CarName == "" || requestBody.Company == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.CarErrorResponse(errors.New(
				"Please specify title and author")))
		}

		result, err := service.InsertCarService(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.CarErrorResponse(err))
		}

		return c.JSON(presenters.CarSuccessResponse(result))
	}
}

func UpdateCar(service cars.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Car

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.CarErrorResponse(err))
		}

		result, err := service.UpdateCarService(&requestBody)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.CarErrorResponse(err))
		}

		return c.JSON(presenters.CarSuccessResponse(result))
	}
}

func RemoveCar(service cars.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.DeleteRequest
		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenters.CarErrorResponse(err))
		}

		carId := requestBody.ID
		err = service.RemoveCarService(carId)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.CarErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "updated succesfully",
			"err":    nil,
		})
	}
}

func GetCars(service cars.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.CheckCarService()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenters.CarErrorResponse(err))
		}
		return c.JSON(presenters.CarsSuccessResponse(fetched))
	}
}
