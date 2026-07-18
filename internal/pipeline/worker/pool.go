package worker

import (
	"context"
	"sync"
	"time"
)

type WorkerPool[J any, R any] struct {
	poolSize int
}

func NewWorkerPool[J any, R any](poolSize int) *WorkerPool[J, R] {
	return &WorkerPool[J, R]{poolSize: poolSize}
}

func (w *WorkerPool[J, R]) Run(ctx context.Context, Job <-chan J, fn func(context.Context, J) (R, error)) <-chan Result[R] {
	var wg sync.WaitGroup

	ResultStream := make(chan Result[R])
	for workerID := range w.poolSize {
		wg.Add(1)
		go func(ID int) {

			defer wg.Done()

		loop:
			for {
				select {
				case <-ctx.Done():
					break loop
				case jobItem, ok := <-Job:
					if !ok {
						break loop
					}
					startTime := time.Now()
					result, err := fn(ctx, jobItem)
					if err != nil {
						duration := time.Since(startTime)
						ResultStream <- Result[R]{
							Error:    err,
							WorkerID: ID,
							Result:   result,
							Duration: duration,
						}
						continue
					}

					duration := time.Since(startTime)

					ResultStream <- Result[R]{
						WorkerID: ID,
						Result:   result,
						Error:    nil,
						Duration: duration,
					}

				}
			}

		}(workerID)
	}

	go func() {
		wg.Wait()
		close(ResultStream)
	}()
	return ResultStream

}
