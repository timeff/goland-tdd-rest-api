package controller

import (
	"golang-tdd-rest-api/todo"
	"golang-tdd-rest-api/todo/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodoAndReturnZero(t *testing.T) {
	mockTodoRepo := mocks.MockTodoRepo{}
	mockTodoEmpty := []*todo.Todo{}

	mockTodoRepo.On("Get").Return(mockTodoEmpty, nil)
	controller := NewTodoController(&mockTodoRepo)

	todos, err := controller.Get()
	assert.Nil(t, err)
	assert.NotNil(t, todos)
	assert.Len(t, todos, 0)
}

func TestGetTodo(t *testing.T) {
	mockTodoRepo := mocks.MockTodoRepo{}
	mockTodo := []*todo.Todo{
		&todo.Todo{
			ID:      1,
			Content: "Test1",
			Done:    false,
		},
	}

	mockTodoRepo.On("Get").Return(mockTodo, nil)

	controller := NewTodoController(&mockTodoRepo)

	todos, err := controller.Get()
	assert.Nil(t, err)
	assert.NotNil(t, todos)
	assert.Len(t, todos, len(mockTodo))
}

func TestCreateTodo(t *testing.T) {
	mockTodoRepo := mocks.MockTodoRepo{}

	mockID := int64(1)
	mockTodo := todo.Todo{
		ID:      mockID,
		Content: "Test1",
		Done:    false,
	}

	mockTodoRepo.On("Create", &mockTodo).Return(mockID, nil)

	controller := NewTodoController(&mockTodoRepo)

	todoID, err := controller.Create(&mockTodo)
	assert.Nil(t, err)
	assert.Equal(t, mockID, todoID)
}

func TestUpdateTodo(t *testing.T) {
	mockTodoRepo := mocks.MockTodoRepo{}
	mockID := int64(1)
	mockTodo := todo.Todo{
		ID:      mockID,
		Content: "Test1",
		Done:    false,
	}

	mockTodoRepo.On("GetByID", mockID).Return(&mockTodo, nil)
	mockTodoRepo.On("Update", &mockTodo).Return(nil)

	controller := NewTodoController(&mockTodoRepo)

	err := controller.Update(&mockTodo)
	assert.Nil(t, err)
}

func TestUpdateToWrongID(t *testing.T) {
	mockTodoRepo := mocks.MockTodoRepo{}
	mockID := int64(99)
	mockTodo := todo.Todo{
		ID:      mockID,
		Content: "Test1",
		Done:    false,
	}
	expectedError := todo.NOT_FOUND_ERROR

	mockTodoRepo.On("GetByID", mockID).Return(nil, expectedError)

	controller := NewTodoController(&mockTodoRepo)

	err := controller.Update(&mockTodo)
	assert.Equal(t, expectedError, err)
}

func TestDeleteTodo(t *testing.T) {
	mockTodoRepo := mocks.MockTodoRepo{}
	mockID := int64(1)

	mockTodoRepo.On("Delete", mockID).Return(nil)

	controller := NewTodoController(&mockTodoRepo)

	err := controller.Delete(mockID)
	assert.Nil(t, err)
}
