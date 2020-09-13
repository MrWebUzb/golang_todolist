package entities

import (
	"database/sql"
	"fmt"

	"github.com/MrWebUzb/apiserver/app/utils"
)

// TodoEntity ...
type TodoEntity struct {
	db        *sql.DB
	tableName string

	ID        int64           `json:"id"`
	Message   string          `json:"message"`
	IsDone    bool            `json:"is_done"`
	CreatedAt utils.Timestamp `json:"created_at"`
	UpdatedAt utils.Timestamp `json:"updated_at"`
	DeletedAt utils.Timestamp `json:"-"`
}

// NewTodoEntity ...
func NewTodoEntity(db *sql.DB) *TodoEntity {
	return &TodoEntity{
		db:        db,
		tableName: "todo",
	}
}

// GetAll ...
func (te *TodoEntity) GetAll() ([]*TodoEntity, error) {
	rows, err := te.db.Query(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE deleted_at is NULL",
			te.tableName,
		),
	)

	if err != nil {
		return nil, err
	}

	result := []*TodoEntity{}

	defer rows.Close()

	for rows.Next() {
		ent := &TodoEntity{}

		err = rows.Scan(
			&ent.ID,
			&ent.Message,
			&ent.IsDone,
			&ent.CreatedAt,
			&ent.UpdatedAt,
			&ent.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, ent)
	}

	return result, nil
}

// FindByID ...
func (te *TodoEntity) FindByID(id int64) (*TodoEntity, error) {
	stmt, err := te.db.Prepare(
		fmt.Sprintf(
			"SELECT * FROM %s WHERE id=? AND deleted_at is NULL",
			te.tableName,
		),
	)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result := &TodoEntity{}

	err = stmt.QueryRow(id).Scan(
		&result.ID,
		&result.Message,
		&result.IsDone,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
