package user

import (
	"demo/app/models"
)

type Repository interface {
	Create(user *models.User) (int64, error)
	GetByPhoneNumber(id int64) (models.User, error)
}
