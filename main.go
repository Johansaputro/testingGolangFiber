package main

import (
	"context"
	"fmt"
	"log"
	"testingfiber/api/routes"
	"testingfiber/pkg/cars"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, cancel, err := databaseConnection()

	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}

	fmt.Println("Database connection success!")

	carCollection := db.Collection("cars")
	carRepo := cars.NewRepo(carCollection)
	carService := cars.NewService(carRepo)

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo car shop!"))
	})

	api := app.Group("/api")
	routes.CarRouter(api, carService)
	defer cancel()
	log.Fatal(app.Listen(":8080"))

}

func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb://localhost:27017/cars").SetServerSelectionTimeout(5*time.
		Second))
	if err != nil {
		cancel()
		return nil, nil, err
	}
	db := client.Database("cars")
	return db, cancel, nil
}
