package usecase

import (
	"go-clean-architecture-demo/app/entities"

	"github.com/jinzhu/gorm"
	errors "github.com/pkg/errors"
)

func (useCase *useCase) UpdateById(userId uint, userEntity *entities.User) (*entities.User, error) {
	errs := validate.Struct(userEntity)
	if errs != nil {
		return nil, errors.Wrap(errs, "update user by id fail")
	}

	condition := &entities.User{
		Model: gorm.Model{ID: userId},
	}

	_, err := useCase.repo.Update(condition, userEntity)
	userEntity.ID = userId
	return userEntity, err
}
