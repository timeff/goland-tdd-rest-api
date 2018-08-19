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

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

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

func TestGetTodo(t *testing.T) {
	mockTodoController := mocks.MockTodoController{}
	mockTodo := []*todo.Todo{
		&todo.Todo{
			ID:      1,
			Content: "Test1",
			Done:    false,
		},
	}
	mockTodoController.On("Get").Return(mockTodo, nil)

	expectedJSON, _ := json.Marshal(mockTodo)
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
	mockTodoController.On("Create", &mockTodo).Return(mockID, nil)

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
	mockTodoController.On("Update", &mockTodo).Return(nil)

	requestData, err := json.Marshal(&mockTodo)
	assert.NoError(t, err)

	c, rec := createEchoContext(t, echo.PUT, "/todo/", strings.NewReader(string(requestData)))

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
	mockIDString := strconv.Itoa(int(mockID))
	mockTodoController.On("Delete", mockID).Return(nil)

	c, rec := createEchoContext(t, echo.DELETE, "/todo/"+mockIDString, nil)
	c.SetPath("todo/:id")
	c.SetParamNames("id")
	c.SetParamValues(mockIDString)

	handler := HTTPTodoHandler{
		todoController: &mockTodoController,
	}
	handler.Delete(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockTodoController.AssertExpectations(t)
}
