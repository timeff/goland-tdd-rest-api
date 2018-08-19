package repository

import (
	"database/sql"
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
	allTodo := []*todo.Todo{}
	err := g.Conn.Find(&allTodo).Error

	if err == sql.ErrNoRows {
		return allTodo, nil
	}
	if err != nil {
		return nil, err
	}

	return allTodo, nil
}

func (g *GormTodoRepo) Create(t *todo.Todo) (int64, error) {
	db := g.Conn.Create(t)
	if err := db.Error; err != nil {
		return 0, err
	}

	lastInsertID := t.ID

	return lastInsertID, nil
}

func (g *GormTodoRepo) Update(t *todo.Todo) error {
	err := g.Conn.Save(t).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *GormTodoRepo) Delete(id int64) error {
	mockTodo := todo.Todo{
		ID: id,
	}

	err := g.Conn.Delete(&mockTodo).Error
	if err != nil {
		return err
	}

	return nil
}
