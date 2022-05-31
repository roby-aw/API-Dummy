package dummy

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindFoodByName(name string) (foods []Food, err error)
}

type Service interface {
	GetFoodByName(name string) (foods []Food, err error)
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

func (s *service) GetFoodByName(name string) (foods []Food, err error) {
	fmt.Println("service jalan")
	return s.repository.FindFoodByName(name)
}
