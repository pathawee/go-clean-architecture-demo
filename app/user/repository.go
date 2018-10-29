package user

import (
	"go-clean-architecture-demo/app/entities"
)

type Repository interface {
	Create(user *entities.User) (int64, error)
	Update(condition *entities.User, user *entities.User) (int64, error)
	List(skip int, limit int) ([]*entities.User, error)
	GetByPhoneNumber(id int64) (entities.User, error)
}
