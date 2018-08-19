package repository

import (
	"database/sql"
	"golang-tdd-rest-api/todo"

	"github.com/jinzhu/gorm"
)

type GormTodoRepo struct {
	conn *gorm.DB
}

func NewGormRepo(Conn *gorm.DB) todo.Repository {
	return &GormTodoRepo{
		conn: Conn,
	}
}

func (g *GormTodoRepo) Get() ([]*todo.Todo, error) {
	allTodo := []*todo.Todo{}
	err := g.conn.Find(&allTodo).Error

	if err == sql.ErrNoRows {
		return allTodo, nil
	}
	if err != nil {
		return nil, err
	}

	return allTodo, nil
}

func (g *GormTodoRepo) GetByID(id int64) (*todo.Todo, error) {
	mockTodo := todo.Todo{
		ID: id,
	}

	err := g.conn.First(&mockTodo).Error
	if err == sql.ErrNoRows {
		return nil, todo.NOT_FOUND_ERROR
	}
	if err != nil {
		return nil, err
	}

	return &mockTodo, nil
}

func (g *GormTodoRepo) Create(t *todo.Todo) (int64, error) {
	db := g.conn.Create(t)
	if err := db.Error; err != nil {
		return 0, err
	}

	lastInsertID := t.ID

	return lastInsertID, nil
}

func (g *GormTodoRepo) Update(t *todo.Todo) error {
	err := g.conn.Model(t).Update(map[string]interface{}{
		"content": t.Content,
		"done":    t.Done,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (g *GormTodoRepo) Delete(id int64) error {
	mockTodo := todo.Todo{
		ID: id,
	}

	err := g.conn.Delete(&mockTodo).Error
	if err != nil {
		return err
	}

	return nil
}
