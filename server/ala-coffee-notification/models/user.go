package models

import (
	"database/sql/driver"
	"errors"
	"time"
)

type User struct {
	ID          string     `json:"id" db:"id"`
	Email       string     `json:"email" db:"email"`
	Username    *string    `json:"username" db:"username"`
	Password    string     `json:"password" db:"password"`
	FirstName   string     `json:"firstName" db:"first_name"`
	LastName    string     `json:"lastName" db:"last_name"`
	DateOfBirth time.Time  `json:"dateOfBirth" db:"date_of_birth"`
	ImageURL    *string    `json:"imageUrl" db:"image_url"`
	Story       *string    `json:"story" db:"story"`
	CreatedAt   *time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt   *time.Time `json:"deletedAt" db:"deleted_at"`
	Status      UserStatus `json:"status" db:"status"`
}

type SignUpAccount struct {
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
	DateOfBirth string `json:"dateOfBirth" validate:"required,date_string"`
}

type SignInAccount struct {
	Password   string `json:"password" validate:"required"`
	Credential string `json:"credential" validate:"required"`
}

type UserStatus string

const (
	UserStatusInactive UserStatus = "inactive"
	UserStatusActive   UserStatus = "active"
)

func (val *UserStatus) Value() (driver.Value, error) {
	return string(*val), nil
}

func (val *UserStatus) Scan(value interface{}) error {
	switch s := value.(type) {
	case string:
		*val = UserStatus(s)
	case []byte:
		*val = UserStatus(string(s))
	default:
		return errors.New("incompatible type for UserStatus")
	}

	return nil
}
