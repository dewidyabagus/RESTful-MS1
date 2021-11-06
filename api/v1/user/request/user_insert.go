package request

import "RESTfulMS1/business/user"

type RequestUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}

func (r *RequestUser) ToUserAddSpec() *user.UserAddSpec {
	return &user.UserAddSpec{
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Role:      r.Role,
		Password:  r.Password,
	}
}
