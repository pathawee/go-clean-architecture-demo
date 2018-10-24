package post

import (
	"demo/app/models"
)

type UseCase interface {
	Create(*models.Post) (*models.Post, error)
}
