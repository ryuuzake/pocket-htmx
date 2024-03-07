package model

import "github.com/pocketbase/pocketbase/models"

var _ models.Model = (*Todo)(nil)

const (
	TODOS = "todos"
)

type Todo struct {
	models.BaseModel

	ID         string `db:"id"`
	Content    string `db:"content"`
	IsComplete bool   `db:"is_complete"`
}

func (*Todo) TableName() string {
	return TODOS
}
