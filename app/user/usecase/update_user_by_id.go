package usecase

import (
	"go-clean-architecture-demo/app/entities"
)

func (useCase *useCase) UpdateById(userId int64, userEntity *entities.User) (*entities.User, error) {
’”	condition := &entities.User{ID: userId}
	_, err := useCase.repo.Update(condition, userEntity)
	userEntity.ID = userId
	return userEntity, err
}
