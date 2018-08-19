package todo

import (
	"errors"
	"time"
)

type Todo struct {
	ID        int64      `json:"id`
	Content   string     `json:"content"`
	Done      bool       `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var (
	NOT_FOUND_ERROR = errors.New("No todo for given ID")
	WRONG_PARAMS    = errors.New("Wrong params")
)
