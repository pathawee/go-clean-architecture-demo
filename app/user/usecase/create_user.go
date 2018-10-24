package usecase

import (
	"go-clean-architecture-demo/app/entities"
)

func (useCase *useCase) Create(userEntity *entities.User) (*entities.User, error) {
	useCase.repo.Create(userEntity)
	return userEntity, nil
}
