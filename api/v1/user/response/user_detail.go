package response

import "RESTfulMS1/business/user"

type UserDetail struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

func GetUserResponseDetail(user *user.User) *UserDetail {
	return &UserDetail{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
}
