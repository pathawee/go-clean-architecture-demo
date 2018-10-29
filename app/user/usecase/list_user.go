package usecase

import (
	"go-clean-architecture-demo/app/entities"
)

func (useCase *useCase) List(skip int, limit int) ([]*entities.User, error) {
	if skip < 0 {
		skip = 0
	}

	if limit < 1 {
		limit = 1
	}

	if limit > 50 {
		limit = 50
	}

	return useCase.repo.List(skip, limit)
}
