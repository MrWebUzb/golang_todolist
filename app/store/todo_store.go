package store

import (
	"database/sql"

	"github.com/MrWebUzb/apiserver/app/entities"
)

// TodoStore ...
type TodoStore struct {
	entity *entities.TodoEntity
}

// NewTodoStore ...
func NewTodoStore(db *sql.DB) *TodoStore {
	return &TodoStore{
		entity: entities.NewTodoEntity(db),
	}
}

// GetAll ...
func (ts *TodoStore) GetAll() ([]*entities.TodoEntity, error) {
	return ts.entity.GetAll()
}

// FindByID ...
func (ts *TodoStore) FindByID(id int64) (*entities.TodoEntity, error) {
	return ts.entity.FindByID(id)
}
