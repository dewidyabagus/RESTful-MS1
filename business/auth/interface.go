package auth

type Service interface {
	UserLoginWithEmailPassword(user *LoginSpec) (token *string, err error)
}
