package main

import (
	"go-fiber-project/app"
	"go-fiber-project/configs"
	"go-fiber-project/repository"
	"go-fiber-project/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoReportsitoryDb(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodos)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)

	appRoute.Listen(":8080")
}
