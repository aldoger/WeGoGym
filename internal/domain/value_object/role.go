package valueobject

import (
	"database/sql/driver"
	"fmt"
)

type UserRole struct {
	value string
}

const (
	MEMBER_ROLE = "member"
	ADMIN_ROLE  = "admin"
)

func NewUserRole(value string) UserRole {
	if value != MEMBER_ROLE && value != ADMIN_ROLE {
		return UserRole{}
	}
	return UserRole{value: value}
}

func (r UserRole) Value() driver.Value {
	return r.value
}

func (r *UserRole) GetRole() string {
	return r.value
}

func (r UserRole) IsAdmin(role string) bool {
	return role == ADMIN_ROLE
}

func (r *UserRole) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("expected string for UserRole but got %T", value)
	}
	role := NewUserRole(strVal)
	*r = role
	return nil
}
