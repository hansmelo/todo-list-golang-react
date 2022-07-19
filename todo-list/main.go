package main

import (
	"com/hans/todolist/database"
	"com/hans/todolist/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "host=localhost port=5432 user=keycloak-postgres password=keycloak-postgres dbname=todolist sslmode=disable"
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Connected to database!")
	database.DBConn.AutoMigrate(&models.Todo{})
	fmt.Println("Database migrated!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", models.GetTodos)
	app.Get("/todos/:id", models.GetById)
	app.Post("/todos", models.CreateTodos)
	app.Put("/todos/:id", models.UpdateTodos)
	app.Delete("/todos/:id", models.DeleteTodos)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initDatabase()
	setupRoutes(app)
	app.Listen(":8888")
}
