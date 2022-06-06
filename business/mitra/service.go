package mitra

import "github.com/go-playground/validator/v10"

type Repository interface {
}

type Service interface {
}

type service struct {
	repository Repository
	validate   *validator.Validate
}
