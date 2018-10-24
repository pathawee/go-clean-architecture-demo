package post

import (
	"demo/app/models"
)

type Repository interface {
	Create(*models.Post) (int64, error)
}
