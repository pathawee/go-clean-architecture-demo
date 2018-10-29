package user

import (
	"go-clean-architecture-demo/app/entities"
)

type Repository interface {
	Create(user *entities.User) (uint, error)
	Update(condition *entities.User, user *entities.User) (uint, error)
	List(skip int, limit int) ([]*entities.User, error)
	GetByPhoneNumber(id int64) (entities.User, error)
}
