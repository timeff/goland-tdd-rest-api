package http

import (
	"encoding/json"
	"golang-tdd-rest-api/todo"
	"golang-tdd-rest-api/todo/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func createEchoContext(t *testing.T,
	method string,
	url string,
	body io.Reader,
) (
	echo.Context,
	*httptest.ResponseRecorder,
) {
	e := echo.New()
	req, err := http.NewRequest(method, url, body)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	return e.NewContext(req, rec), rec
}

func TestGetTodoAndReturnEmpty(t *testing.T) {
	mockTodoController := mocks.MockTodoController{}
	mockTodoEmpty := []*todo.Todo{}
	mockTodoController.On("Get").Return(mockTodoEmpty, nil)

	expectedJSON, _ := json.Marshal([]interface{}{})
	expectedJSONString := string(expectedJSON)

	c, rec := createEchoContext(t, echo.GET, "/todo", nil)

	handler := HTTPTodoHandler{
		todoController: &mockTodoController,
	}
	handler.Get(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expectedJSONString, rec.Body.String())
	mockTodoController.AssertExpectations(t)
}

func TestCreateTodo(t *testing.T) {
	mockTodoController := mocks.MockTodoController{}
	mockID := int64(1)
	mockTodo := todo.Todo{
		ID:      mockID,
		Content: "Test1",
		Done:    false,
	}
	mockTodoController.On("Create").Return(mockID, nil)

	requestData, err := json.Marshal(&mockTodo)
	assert.NoError(t, err)

	expectedJSON, err := json.Marshal(&JSONCreateReturn{
		ID: mockID,
	})
	assert.NoError(t, err)
	expectedJSONString := string(expectedJSON)

	c, rec := createEchoContext(t, echo.POST, "/todo", strings.NewReader(string(requestData)))

	handler := HTTPTodoHandler{
		todoController: &mockTodoController,
	}
	handler.Create(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expectedJSONString, rec.Body.String())
	mockTodoController.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	mockTodoController := mocks.MockTodoController{}
	mockID := int64(1)
	mockTodo := todo.Todo{
		ID:      mockID,
		Content: "Test1",
		Done:    false,
	}
	mockTodoController.On("Update").Return(nil)

	requestData, err := json.Marshal(&mockTodo)
	assert.NoError(t, err)

	c, rec := createEchoContext(t, echo.PUT, "/todo/"+strconv.Itoa(int(mockID)), strings.NewReader(string(requestData)))

	handler := HTTPTodoHandler{
		todoController: &mockTodoController,
	}
	handler.Update(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockTodoController.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	mockTodoController := mocks.MockTodoController{}
	mockID := int64(1)
	mockTodoController.On("Delete").Return(nil)

	c, rec := createEchoContext(t, echo.DELETE, "/todo/"+strconv.Itoa(int(mockID)), nil)

	handler := HTTPTodoHandler{
		todoController: &mockTodoController,
	}
	handler.Delete(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockTodoController.AssertExpectations(t)
}
