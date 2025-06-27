package valueobject

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type UserRole struct {
	value string
}

const (
	MEMBER_ROLE = "member"
	ADMIN_ROLE  = "admin"
)

func NewUserRole(value string) (UserRole, error) {
	if value != MEMBER_ROLE && value != ADMIN_ROLE {
		return UserRole{}, errors.New("invalid user role")
	}
	return UserRole{value: value}, nil
}

func (r UserRole) Value() (driver.Value, error) {
	return r.value, nil
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
	role, err := NewUserRole(strVal)
	if err != nil {
		return err
	}
	*r = role
	return nil
}
