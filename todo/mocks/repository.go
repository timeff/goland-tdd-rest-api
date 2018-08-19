package mocks

import (
	"golang-tdd-rest-api/todo"

	"github.com/stretchr/testify/mock"
)

type MockTodoRepo struct {
	mock.Mock
}

func (m *MockTodoRepo) Get() ([]*todo.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*todo.Todo), args.Error(1)
}

func (m *MockTodoRepo) GetByID(id int64) (*todo.Todo, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*todo.Todo), args.Error(1)
}

func (m *MockTodoRepo) Create(t *todo.Todo) (int64, error) {
	args := m.Called(t)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTodoRepo) Update(t *todo.Todo) error {
	args := m.Called(t)
	return args.Error(0)
}

func (m *MockTodoRepo) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
