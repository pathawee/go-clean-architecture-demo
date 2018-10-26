package usecase

import (
	"go-clean-architecture-demo/app/entities"

	errors "github.com/pkg/errors"
)

func (useCase *useCase) UpdateById(userId int64, userEntity *entities.User) (*entities.User, error) {
	errs := validate.Struct(userEntity)
	if errs != nil {
		return nil, errors.Wrap(errs, "update user by id fail")
	}

	condition := &entities.User{ID: userId}

	_, err := useCase.repo.Update(condition, userEntity)
	userEntity.ID = userId
	return userEntity, err
}
