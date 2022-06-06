package mitra

import (
	"github.com/go-playground/validator/v10"
)

type Repository interface {
}

type Service interface {
	PostInputPoin(input InputPoinMitra) error
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) PostInputPoin(data InputPoinMitra) error {
	var err error
	return err
}
