package user

import (
	"demo/app/models"
)

type UseCase interface {
	Create(*models.User) (*models.User, error)
}
