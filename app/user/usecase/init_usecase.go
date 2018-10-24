package usecase

import (
	"go-clean-architecture-demo/app/user"
)

type useCase struct {
	repo user.Repository
}

func InitUseCase(repo user.Repository) user.UseCase {
	return &useCase{
		repo: repo,
	}
}
