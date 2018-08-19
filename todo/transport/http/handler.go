package http

import (
	"errors"
	"golang-tdd-rest-api/todo"

	"github.com/labstack/echo"
)

type HTTPTodoHandler struct {
	todoController todo.Controller
}

func NewHTTPTodoHandler(e *echo.Echo, todoController todo.Controller) {
	handler := &HTTPTodoHandler{
		todoController,
	}

	e.GET("/todo", handler.Get)
	e.POST("/todo", handler.Create)
	e.PUT("/todo/:id", handler.Update)
	e.DELETE("/todo/:id", handler.Delete)
}

func (a *HTTPTodoHandler) Get(c echo.Context) error {
	return errors.New("Need implement")
}

type JSONCreateReturn struct {
	ID int64 `json:"id"`
}

func (a *HTTPTodoHandler) Create(c echo.Context) error {
	return errors.New("Need implement")
}

func (a *HTTPTodoHandler) Update(c echo.Context) error {
	return errors.New("Need implement")
}

func (a *HTTPTodoHandler) Delete(c echo.Context) error {
	return errors.New("Need implement")
}
