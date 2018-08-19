package main

import (
	"golang-tdd-rest-api/database"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	todoController "golang-tdd-rest-api/todo/controller"
	todoRepo "golang-tdd-rest-api/todo/repository"
	todoHTTPTransport "golang-tdd-rest-api/todo/transport/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("unable to read .env file")
	}
}

func main() {
	e := echo.New()

	gormTodo, err := database.GetGormTodo()
	if err != nil {
		log.Fatal(err)
	}

	todoRepo := todoRepo.NewGormRepo(gormTodo)
	todoController := todoController.NewTodoController(todoRepo)
	todoHTTPTransport.NewHTTPTodoHandler(e, todoController)

	port := os.Getenv("WEB_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(e.Start(":" + port))
}