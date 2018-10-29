package user

import (
	"go-clean-architecture-demo/app/entities"
)

type UseCase interface {
	Create(*entities.User) (*entities.User, error)
	UpdateById(int64, *entities.User) (*entities.User, error)
	List(skip int, limit int) ([]*entities.User, error)
}
