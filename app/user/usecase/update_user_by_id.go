package usecase

import (
	"go-clean-architecture-demo/app/entities"
)

func (useCase *useCase) UpdateById(userId int64, userEntity *entities.User) (*entities.User, error) {
	// condition := map[string]interface{}{"ID": userId}
	condition := &entities.User{ID: userId}
	useCase.repo.Update(condition, userEntity)
	return userEntity, nil
}
