package user

import (
	"go-clean-architecture-demo/app/entities"
)

type Repository interface {
	Create(user *entities.User) (int64, error)
	GetByPhoneNumber(id int64) (entities.User, error)
}
