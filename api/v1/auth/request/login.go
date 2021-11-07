package request

import (
	"RESTfulMS1/business/auth"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) ToLoginSpec() *auth.LoginSpec {
	return &auth.LoginSpec{
		Email:    l.Email,
		Password: l.Password,
	}
}
