package todo

import (
	"errors"
	"time"
)

// A todo is list of todo
// swagger:response todo
type Todo struct {
	// the id for this todo
	//
	// unique: true
	ID int64 `json:"id`
	// the content for this todo
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
