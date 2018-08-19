package repository

import (
	"golang-tdd-rest-api/todo"

	"github.com/jinzhu/gorm"
)

type GormTodoRepo struct {
	Conn *gorm.DB
}

func NewGormRepo(Conn *gorm.DB) todo.Repository {
	return &GormTodoRepo{Conn}
}

func (g *GormTodoRepo) Get() ([]*todo.Todo, error) {
	return nil, nil
}

func (g *GormTodoRepo) Create(*todo.Todo) (int64, error) {
	return 0, nil
}

func (g *GormTodoRepo) Update(*todo.Todo) (int64, error) {
	return 0, nil
}

func (g *GormTodoRepo) Delete(id int64) error {
	return nil
}
