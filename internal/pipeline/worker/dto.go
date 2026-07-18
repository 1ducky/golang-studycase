package worker

import "time"

type Result[T any] struct {
	Result   T
	Error    error
	WorkerID int
	Duration time.Duration
}

type WorkerStatus string

const (
	WorkerActive  WorkerStatus = "active"
	WorkerSuccess WorkerStatus = "success"
	WorkerFailed  WorkerStatus = "failed"
	WorkerDie     WorkerStatus = "die"
)
