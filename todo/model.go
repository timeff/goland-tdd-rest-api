package todo

import (
	"time"
)

type Todo struct {
	ID        int64     `json:"id`
	Content   string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
