package models

import (
	"com/hans/todolist/database"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        uint   `gorm:primarykey" json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todos []Todo
	db.Order("title asc").Find(&todos)
	return c.JSON(todos)
}

func GetById(c *fiber.Ctx) error {
	db := database.DBConn
	var todo Todo
	id := c.Params("id")
	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find an item", "error": err.Error()})
	}
	return c.JSON(todo)
}

func CreateTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check your input", "error": err.Error()})
	}
	if err := db.Create(&todo).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create an item", "error": err.Error()})
	}
	return c.JSON(todo)
}

func UpdateTodos(c *fiber.Ctx) error {
	type TodoUpdate struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	db := database.DBConn
	var todo Todo
	id := c.Params("id")
	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find an item", "error": err.Error()})
	}

	var todoUpdate TodoUpdate
	if err := c.BodyParser(&todoUpdate); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check your input", "error": err.Error()})
	}

	todo.Title = todoUpdate.Title
	todo.Completed = todoUpdate.Completed
	if err := db.Save(&todo).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update an item", "error": err.Error()})
	}

	return c.JSON(todo)
}

func DeleteTodos(c *fiber.Ctx) error {
	db := database.DBConn
	var todo Todo

	id := c.Params("id")
	if err := db.First(&todo, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find an item", "error": err.Error()})
	}

	if err := db.Delete(&todo).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete an item", "error": err.Error()})
	}

	return c.SendStatus(204)
}
