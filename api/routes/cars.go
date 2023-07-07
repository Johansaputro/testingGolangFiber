package routes

import (
	"testingfiber/api/handlers"
	"testingfiber/pkg/cars"

	"github.com/gofiber/fiber/v2"
)

func CarRouter(app fiber.Router, service cars.Service) {
	app.Get("/cars", handlers.GetCars(service))
	app.Post("/cars", handlers.AddCar(service))
	app.Put("/cars", handlers.UpdateCar(service))
	app.Delete("/cars", handlers.RemoveCar(service))
}
