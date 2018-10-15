package http

import (
	"golang-tdd-rest-api/todo"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type HTTPTodoHandler struct {
	todoController todo.Controller
}

func NewHandler(e *echo.Echo, todoController todo.Controller) {
	handler := &HTTPTodoHandler{
		todoController,
	}
	// swagger:route GET /todo todo Get
	//
	// Lists all todo.
	//
	// This will show all available todo.
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       200: todo
	e.GET("/todo", handler.Get)

	e.POST("/todo", handler.Create)
	e.PUT("/todo", handler.Update)
	e.DELETE("/todo/:id", handler.Delete)
}

func (h *HTTPTodoHandler) Get(c echo.Context) error {
	todos, err := h.todoController.Get()
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}

	return c.JSON(http.StatusOK, todos)
}

type JSONCreateReturn struct {
	ID int64 `json:"id"`
}

func (h *HTTPTodoHandler) Create(c echo.Context) error {
	todo := todo.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}
	insertID, err := h.todoController.Create(&todo)
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}

	result := &JSONCreateReturn{
		ID: insertID,
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HTTPTodoHandler) Update(c echo.Context) error {
	todo := todo.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}

	err = h.todoController.Update(&todo)
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}

	return c.NoContent(http.StatusOK)
}

func (h *HTTPTodoHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		errorHTTPCode, errorReponse := getErrorResponse(todo.WRONG_PARAMS)
		return c.JSON(errorHTTPCode, errorReponse)
	}

	idP, err := strconv.Atoi(id)
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}
	id64 := int64(idP)

	err = h.todoController.Delete(id64)
	if err != nil {
		errorHTTPCode, errorReponse := getErrorResponse(err)
		return c.JSON(errorHTTPCode, errorReponse)
	}

	return c.NoContent(http.StatusOK)
}
