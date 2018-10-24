package repository

import (
	"demo/app/models"
	"demo/app/user"
	"github.com/jinzhu/gorm"
	"log"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func InitMysqlRepository(Conn *gorm.DB) user.Repository {
	return &mysqlRepository{Conn}
}

func (mysqlRepo *mysqlRepository) Create(userEntity *models.User) (int64, error) {
	if err := mysqlRepo.Conn.Save(&userEntity).Error; err != nil {
		log.Print("Saving error: ", err)
	}

	return userEntity.ID, nil
}

func (mysqlRepo *mysqlRepository) GetByPhoneNumber(id int64) (models.User, error) {
	userData := models.User{}

	if err := mysqlRepo.Conn.First(&userData, id).Error; err != nil {
		log.Print("Saving error: ", err)
	}

	return userData, nil
}
