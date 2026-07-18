package logging

import (
	"context"
	"crypto/rand"
	"fmt"
)

type loggingContext struct{}

var loggingKey loggingContext

func WithRequestID(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, loggingKey, requestId)
}

func GetRequestID(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(loggingKey).(string)
	fmt.Printf("Getting auth: %+v, ok: %v\n", user, ok)
	return user, ok
}

func GenerateRandomID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
