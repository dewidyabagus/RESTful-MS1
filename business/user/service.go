package user

import (
	"RESTfulMS1/utils/validator"

	"RESTfulMS1/business"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) AddNewUser(user *UserAddSpec) error {
	err := validator.GetValidator().Struct(user)
	if err != nil {
		return business.ErrDataNotSpec
	}

	return s.repository.AddNewUser(user.toInsertUser())
}

func (s *service) GetUserWithEmailPassword(email *string, password *string) (id *string, err error) {
	return s.repository.GetUserWithEmailPassword(email, password)
}

func (s *service) FindUserByUserId(id *string) (*User, error) {
	return s.repository.FindUserByUserId(id)
}
