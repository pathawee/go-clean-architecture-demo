package post

import (
	"go-clean-architecture-demo/app/entities"
)

type UseCase interface {
	Create(*entities.Post) (*entities.Post, error)
}
