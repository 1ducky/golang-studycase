package todos

import (
	"context"
	"database/sql"
	"errors"
	"restApi/internal/database"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Todo, error)
	GetById(ctx context.Context, id int) (Todo, error)
	Create(ctx context.Context, payload CreateRequest) (int, error)
	Update(ctx context.Context, payload UpdateRequest) (int, error)
	Delete(ctx context.Context, payload DeleteRequest) error
}

type TodoMemory struct {
	db database.DBTX
}

func NewTodoMemory(Db database.DBTX) Repository {
	return &TodoMemory{db: Db}
}

func (r *TodoMemory) GetAll(ctx context.Context) ([]Todo, error) {
	query := "SELECT id,task,is_done,created_at,updated_at FROM todos"
	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := make([]Todo, 0, 100)

	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Task,
			&todo.IsDone,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoMemory) GetById(ctx context.Context, id int) (Todo, error) {
	query := "SELECT id,task,is_done,created_at,updated_at FROM todos where id = ?"
	rows := r.db.QueryRowContext(ctx, query, id)

	var todo Todo
	err := rows.Scan(
		&todo.ID,
		&todo.Task,
		&todo.IsDone,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Todo{}, ErrTodoNotFound
		}
		return todo, err
	}
	return todo, nil
}

func (r *TodoMemory) Create(ctx context.Context, payload CreateRequest) (int, error) {
	query := "INSERT INTO todos(task,is_done) VALUES(?,?)"
	result, err := r.db.ExecContext(ctx, query, payload.Task, false)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), nil
}

func (r *TodoMemory) Update(ctx context.Context, payload UpdateRequest) (int, error) {
	query := "UPDATE todos SET task = ?, is_done = ? WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, payload.Task, payload.IsDone, payload.ID)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return 0, err
	}

	return payload.ID, nil
}

func (r *TodoMemory) Delete(ctx context.Context, payload DeleteRequest) error {
	query := "DELETE FROM todos where id = ?"
	result, err := r.db.ExecContext(ctx, query, payload.ID)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrTodoNotFound
	}

	return nil
}

func lastId(todos []Todo) int {
	if len(todos) == 0 {
		return 0
	}
	return todos[len(todos)-1].ID
}
