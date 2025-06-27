package middleware

import valueobject "go-kpl/internal/domain/value_object"

type Middleware struct {
	UserRole *valueobject.UserRole
}

func New() Middleware {
	var UserRole valueobject.UserRole
	return Middleware{
		UserRole: &UserRole,
	}
}
