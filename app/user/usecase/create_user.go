package usecase

import (
	"go-clean-architecture-demo/app/entities"

	"github.com/pkg/errors"
)

func (useCase *useCase) Create(userEntity *entities.User) (*entities.User, error) {
	errs := validate.Struct(userEntity)
	if errs != nil {
		return nil, errors.Wrap(errs, "create user fail")
	}

	useCase.repo.Create(userEntity)
	userEntity.Password = ""
	return userEntity, nil
}
