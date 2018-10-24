package post

import (
	"go-clean-architecture-demo/app/entities"
)

type Repository interface {
	Create(*entities.Post) (int64, error)
}
