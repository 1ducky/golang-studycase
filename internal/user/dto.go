package user

type result[T any] struct {
	data T
	err  error
}

type Role string

var RoleUser Role = "USER"
var RoleAdmin Role = "ADMIN"

type Response struct {
	ok bool
}
