package todolog

import (
	"context"
	"database/sql"
	"errors"
	"restApi/internal/database"
)

type Reposioty interface {
	Create(ctx context.Context, payload CreateRequest) (TodoLog, error)
	GetById(ctx context.Context, id int) (TodoLog, error)
}

type DBRepository struct {
	DB database.DBTX
}

func NewTodoLogRepository(Db database.DBTX) Reposioty {
	return &DBRepository{DB: Db}
}

func (r *DBRepository) Create(ctx context.Context, payload CreateRequest) (TodoLog, error) {
	query := "INSERT INTO `todos_log`( `user_id`, `task_id`, `from_status`, `to_status`) VALUES (?,?,?,?)"
	result, err := r.DB.ExecContext(ctx, query, payload.UserID, payload.TaskID, payload.FromStatus, payload.ToStatus)
	if err != nil {
		return TodoLog{}, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return TodoLog{}, err
	}
	return TodoLog{ID: int(lastId), TaskId: payload.TaskID, UserId: payload.UserID, FromStatus: payload.FromStatus, ToStatus: payload.ToStatus}, nil
}

func (r *DBRepository) GetById(ctx context.Context, id int) (TodoLog, error) {
	query := "SELECT id,user_id,task_id,from_status,to_status FROM todos_log where task_id = ? ORDER BY created_at DESC"
	rows := r.DB.QueryRowContext(ctx, query, id)

	var todoLog TodoLog
	err := rows.Scan(
		&todoLog.ID,
		&todoLog.UserId,
		&todoLog.TaskId,
		&todoLog.FromStatus,
		&todoLog.ToStatus,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TodoLog{}, err
		}
		return TodoLog{}, err
	}
	return todoLog, nil
}
