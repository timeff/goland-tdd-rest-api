package controller

import (
	"golang-tdd-rest-api/todo"
)

type TodoController struct {
	todoRepo todo.Repository
}

func NewTodoController(todoRepo todo.Repository) todo.Controller {
	return &TodoController{
		todoRepo: todoRepo,
	}
}

func (c *TodoController) Get() ([]*todo.Todo, error) {
	todos, err := c.todoRepo.Get()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (c *TodoController) Create(t *todo.Todo) (int64, error) {
	id, err := c.todoRepo.Create(t)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *TodoController) Update(t *todo.Todo) error {
	_, err := c.todoRepo.GetByID(t.ID)
	if err != nil {
		return err
	}

	err = c.todoRepo.Update(t)
	if err != nil {
		return err
	}

	return nil
}

func (c *TodoController) Delete(id int64) error {
	err := c.todoRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
