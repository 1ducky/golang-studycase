package auth

import "context"

func CheckRole(ctx context.Context, role ...Role) (*AuthClaims, bool) {
	user, ok := GetAuht(ctx)
	if !ok {
		return &AuthClaims{}, false
	}
	for _, r := range role {
		if r == user.Role {
			return user, true
		}
	}

	return &AuthClaims{}, false
}
