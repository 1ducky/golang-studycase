package pipeline

import (
	"context"
)

func Batching[T any](ctx context.Context, batchSize int, stream <-chan T) <-chan []T {

	BatchChanel := make(chan []T, 1)
	go func() {
		defer close(BatchChanel)
		batchBuffer := make([]T, 0, batchSize)

		flush := func() bool {
			if len(batchBuffer) == 0 {
				return false
			}
			select {
			case <-ctx.Done():
				return false
			case BatchChanel <- batchBuffer:
				// log.Println("batched", batchBuffer)
				batchBuffer = make([]T, 0, batchSize)
				return true
			}
		}

		for BatchData := range stream {
			select {
			case <-ctx.Done():
				return
			default:
				batchBuffer = append(batchBuffer, BatchData)
				if len(batchBuffer) >= batchSize {
					if !flush() {
						return
					}
				}
			}
		}
		flush()
	}()
	return BatchChanel

}
