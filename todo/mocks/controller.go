package mocks

import (
	"golang-tdd-rest-api/todo"

	"github.com/stretchr/testify/mock"
)

type MockTodoController struct {
	mock.Mock
}

func (m *MockTodoController) Get() ([]*todo.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*todo.Todo), args.Error(1)
}

func (m *MockTodoController) Create(t *todo.Todo) (int64, error) {
	args := m.Called(t)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTodoController) Update(t *todo.Todo) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *MockTodoController) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
