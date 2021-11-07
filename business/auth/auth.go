package auth

type LoginSpec struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
