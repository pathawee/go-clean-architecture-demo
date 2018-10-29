package usecase

import (
	"go-clean-architecture-demo/app/user"

	validator "gopkg.in/go-playground/validator.v8"
)

type useCase struct {
	repo user.Repository
}

var validate *validator.Validate

func InitUseCase(repo user.Repository) user.UseCase {
	// init validator
	config := &validator.Config{TagName: "validate"}
	validate = validator.New(config)

	return &useCase{
		repo: repo,
	}
}
