package auth

import (
	"context"
)

type authContext struct{}

var authKey authContext

func WithAuth(ctx context.Context, u *AuthClaims) context.Context {
	return context.WithValue(ctx, authKey, u)
}

func GetAuht(ctx context.Context) (*AuthClaims, bool) {
	user, ok := ctx.Value(authKey).(*AuthClaims)
	return user, ok
}
