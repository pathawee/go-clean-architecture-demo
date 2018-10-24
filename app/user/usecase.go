package user

import (
	"go-clean-architecture-demo/app/entities"
)

type UseCase interface {
	Create(*entities.User) (*entities.User, error)
}
