package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/gorm"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"golang-tdd-rest-api/todo"
)

func fixedFullRe(s string) string {
	return fmt.Sprintf("^%s$", regexp.QuoteMeta(s))
}

func setUpDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Can't create SQL MOCK: %s", err)
	}

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		log.Fatalf("Can't open GORM connection: %s", err)
	}
	gormDB.LogMode(true)

	return gormDB, mock
}

func createRowsForTodo(todo []todo.Todo) *sqlmock.Rows {
	todoFields := []string{
		"id",
		"content",
		"done",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	rows := sqlmock.NewRows(todoFields)

	for _, t := range todo {
		rows = rows.AddRow(
			t.ID,
			t.Content,
			t.Done,
			t.CreatedAt,
			t.UpdatedAt,
			t.DeletedAt,
		)
	}

	return rows
}

func TestGetTodoAndReturnZero(t *testing.T) {
	db, mock := setUpDB()
	defer db.Close()

	mock.ExpectQuery(fixedFullRe("SELECT * FROM todo;")).
		WillReturnError(sql.ErrNoRows)

	repo := NewGormRepo(db)
	todos, err := repo.Get()

	assert.NoError(t, err)
	assert.NotNil(t, todos)
	assert.Len(t, todos, 0)
}

func TestGetTodo(t *testing.T) {
	db, mock := setUpDB()
	defer db.Close()

	mockTodos := []todo.Todo{
		todo.Todo{
			ID:      int64(1),
			Content: "Test1",
			Done:    false,
		},
	}

	expectedRow := createRowsForTodo(mockTodos)

	mock.ExpectQuery(fixedFullRe("SELECT * FROM todo;")).
		WillReturnRows(expectedRow)

	repo := NewGormRepo(db)
	todos, err := repo.Get()

	assert.NoError(t, err)
	assert.Len(t, todos, len(mockTodos))
}

func TestCreateNewTodo(t *testing.T) {
	db, mock := setUpDB()
	defer db.Close()

	insertID := int64(1)
	content := "Test1"
	done := false

	newTodo := todo.Todo{
		Content: content,
		Done:    done,
	}

	req := fixedFullRe(`
	INSERT INTO "todo" 
		(
			"content",
			"done",
			"created_at",
			"updated_at",
			"deleted_at"
		) 
	VALUES 
		(?,?,?,?,?);`)

	args := []driver.Value{
		content,
		done,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	}

	mock.ExpectExec(req).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(insertID, 1))

	repo := NewGormRepo(db)
	todoID, err := repo.Create(&newTodo)

	assert.NoError(t, err)
	assert.Equal(t, insertID, todoID)
}

func TestUpdateTodo(t *testing.T) {
	db, mock := setUpDB()
	defer db.Close()

	id := int64(1)
	content := "Test1"
	done := false

	updateTodo := todo.Todo{
		ID:      id,
		Content: content,
		Done:    done,
	}

	req := fixedFullRe(`
	UPDATE 
		"todo" 
	SET 
		"content" = ?, 
		"done" = ?,
		"created_at" = ?,
		"updated_at" = ?,
		"deleted_at" = ?
	WHERE 
		"id" = ?;`)
	args := []driver.Value{
		content,
		done,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
		id,
	}

	mock.ExpectExec(req).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(id, 1))

	repo := NewGormRepo(db)
	todoID, err := repo.Update(&updateTodo)

	assert.NoError(t, err)
	assert.Equal(t, id, todoID)
}

func TestDeleteTodo(t *testing.T) {
	db, mock := setUpDB()
	defer db.Close()

	id := int64(1)

	req := fixedFullRe(`
	DELETE FROM
		"todo" 
	WHERE 
		"id" = ?;`)
	args := []driver.Value{
		id,
	}

	mock.ExpectExec(req).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(id, 1))

	repo := NewGormRepo(db)
	err := repo.Delete(id)

	assert.NoError(t, err)
}
