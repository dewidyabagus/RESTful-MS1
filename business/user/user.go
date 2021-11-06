package user

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
	Role      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAddSpec struct {
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Role      string `validate:"required"`
	Password  string `validate:"required"`
}

func (u *UserAddSpec) toInsertUser() *User {
	password := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))

	return &User{
		ID:        uuid.NewString(),
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
