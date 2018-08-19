package database

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gormTodo *gorm.DB

func connGormTodo() (*gorm.DB, error) {
	conn := os.Getenv("MYSQL_CONN")
	if conn == "" {
		return nil, errors.New("env `MYSQL_CONN` not found")
	}

	gormTodo, err := gorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	gormTodo.SingularTable(true)
	return gormTodo, nil
}

func GetGormTodo() (*gorm.DB, error) {
	if gormTodo == nil {
		return connGormTodo()
	}

	if err := gormTodo.DB().Ping(); err != nil {
		return connGormTodo()
	}

	return gormTodo, nil
}

func CloseGormTodo() error {
	if gormTodo != nil {
		if err := gormTodo.Close(); err != nil {
			return err
		}
		gormTodo = nil
	}
	return nil
}
