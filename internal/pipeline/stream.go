package pipeline

import "context"

func Stream[J, R any](ctx context.Context, inputStream <-chan J, buffer int, fn func(context.Context, J) (R, error)) <-chan R {
	outchan := make(chan R, buffer)
	go func() {
		defer close(outchan)
		for job := range inputStream {
			select {
			case <-ctx.Done():
				return
			default:
			}

			res, err := fn(ctx, job)
			if err != nil {
				continue
			}

			select {
			case <-ctx.Done():
				return
			case outchan <- res:
			}

		}
	}()
	return outchan
}
func StreamWithError[J, R any](ctx context.Context, inputStream <-chan J, buffer int, fn func(context.Context, J) (R, error), errFn func(context.Context, any, error)) <-chan R {
	outchan := make(chan R, buffer)
	go func() {
		defer close(outchan)
		for job := range inputStream {
			select {
			case <-ctx.Done():
				return
			default:
			}

			res, err := fn(ctx, job)
			if err != nil {
				errFn(ctx, job, err)
			}

			select {
			case <-ctx.Done():
				return
			case outchan <- res:
			}

		}
	}()
	return outchan
}
