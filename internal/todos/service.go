package todos

import (
	"context"
	"database/sql"
	"io"
	"restApi/internal/auth"
	"restApi/internal/csv"
	"restApi/internal/database"
	"restApi/internal/pipeline"
	"restApi/internal/pipeline/worker"
	"restApi/internal/query"
	todolog "restApi/internal/todos/todo-log"
)

type Service struct {
	repository Repository
	logRepo    todolog.Reposioty
	txManager  *database.TransactionManager
}

func NewService(r Repository, l todolog.Reposioty, tx *database.TransactionManager) *Service {
	return &Service{repository: r, logRepo: l, txManager: tx}
}

func (s *Service) GetAll(ctx context.Context, opt query.GenericQuery) ([]Todo, TodosMeta, error) {
	resource, err := s.repository.GetAll(ctx)
	lenRes := len(resource)
	meta := TodosMeta{
		NextCursor: resource[lenRes-1].ID,
		Limit:      opt.Limit,
		HasNext:    lenRes >= opt.Limit,
	}
	return resource, meta, err
}
func (s *Service) GetById(ctx context.Context, id int) (Todo, error) {
	if id == 0 {
		return Todo{}, ErrTodoBadRequest
	}
	resource, err := s.repository.GetById(ctx, id)
	return resource, err
}
func (s *Service) Create(ctx context.Context, payload CreateRequest) (Todo, error) {
	// auth
	user, ok := auth.GetAuht(ctx)
	if !ok {
		return Todo{}, ErrUnauthorized
	}
	// validation Field
	if payload.Task == "" {
		return Todo{}, ErrTodoBadRequest
	}
	var result Todo

	err := s.txManager.Do(ctx, func(tx *sql.Tx) error {

		todoRepo := NewTodoMemory(tx)
		todoLogRepo := todolog.NewTodoLogRepository(tx)

		id, err := todoRepo.Create(ctx, payload)
		if err != nil {
			return err
		}

		_, err = todoLogRepo.Create(ctx, todolog.CreateRequest{
			UserID:     user.ID,
			TaskID:     id,
			FromStatus: todolog.StatusNull,
			ToStatus:   todolog.StatusCreate,
		})
		if err != nil {
			return err
		}

		resource, err := todoRepo.GetById(ctx, id)
		if err != nil {
			return err
		}

		result = resource // simpan ke variabel luar
		return nil
	})
	if err != nil {
		return Todo{}, err
	}

	return result, nil
}
func (s *Service) Update(ctx context.Context, payload UpdateRequest) (Todo, error) {
	if payload.ID == 0 || payload.Task == "" {
		return Todo{}, ErrTodoBadRequest
	}
	user, ok := auth.GetAuht(ctx)
	if !ok {
		return Todo{}, ErrUnauthorized
	}
	var resource Todo
	err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		todoRepo := NewTodoMemory(tx)
		todoLogRepo := todolog.NewTodoLogRepository(tx)
		var defaultStatus todolog.Status = "CREATE"
		avaliable, err := todoLogRepo.GetById(ctx, payload.ID)
		if err == nil {
			defaultStatus = avaliable.ToStatus
		}

		_, err = todoRepo.Update(ctx, payload)
		if err != nil {
			return err
		}
		_, err = todoLogRepo.Create(ctx, todolog.CreateRequest{UserID: user.ID, TaskID: payload.ID, FromStatus: defaultStatus, ToStatus: todolog.StatusUpdate})
		if err != nil {
			return err
		}
		result, err := todoRepo.GetById(ctx, payload.ID)
		if err != nil {
			return auth.ErrorNotFound
		}
		resource = result

		return nil
	})

	return resource, err
}
func (s *Service) Delete(ctx context.Context, payload DeleteRequest) error {
	if payload.ID == 0 {
		return ErrTodoBadRequest
	}

	user, ok := auth.GetAuht(ctx)
	if !ok {
		return ErrUnauthorized
	}

	err := s.txManager.Do(ctx, func(tx *sql.Tx) error {
		todoRepo := NewTodoMemory(tx)
		todoLogRepo := todolog.NewTodoLogRepository(tx)
		var defaultStatus todolog.Status = "CREATE"
		avaliable, err := todoLogRepo.GetById(ctx, payload.ID)
		if err == nil {
			defaultStatus = avaliable.ToStatus
		}

		err = todoRepo.Delete(ctx, payload)
		if err != nil {
			return err
		}
		_, err = todoLogRepo.Create(ctx, todolog.CreateRequest{UserID: user.ID, TaskID: payload.ID, FromStatus: defaultStatus, ToStatus: todolog.StatusDelete})
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (s *Service) UploadWithCsv(ctx context.Context, payload io.Reader) (BulkResult, error) {
	rowStream := csv.ReadCsvFile(ctx, payload)
	create := pipeline.Stream(ctx, rowStream, 2, func(ctx context.Context, arg []string) (CreateRequest, error) {
		return MapperArgsToTodos(arg)
	})
	batch := pipeline.Batching(ctx, 10, create)
	res, err := s.BulkInsert(ctx, batch, 2)
	return res, err

}

func (s *Service) BulkInsert(ctx context.Context, data <-chan []CreateRequest, workerPool int) (BulkResult, error) {
	var (
		errCount     int64
		successCount int64
		err          error
	)
	pool := worker.NewWorkerPool[[]CreateRequest, []CreateRequest](workerPool)
	report := pool.Run(ctx, data, func(ctx context.Context, cr []CreateRequest) ([]CreateRequest, error) {
		_, err := s.repository.BulkCreate(ctx, cr)
		if err != nil {
			return cr, err
		}
		return cr, nil
	})
	for r := range report {
		if r.Error != nil {
			errCount += int64(len(r.Result))
			err = ErrPartialBulk
		} else {
			successCount += int64(len(r.Result))
		}
	}
	if successCount == 0 {
		err = ErrBulkError
	}

	return BulkResult{SuccessCount: successCount, ErrorCount: errCount, TotalData: successCount + errCount}, err
}
