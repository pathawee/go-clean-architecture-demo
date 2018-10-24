package usecase

import (
	"demo/app/models"
)

func (useCase *useCase) Create(userEntity *models.User) (*models.User, error) {
	useCase.repo.Create(userEntity)
	return userEntity, nil
}
