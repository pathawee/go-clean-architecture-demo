package repository

import (
	"go-clean-architecture-demo/app/entities"
	"go-clean-architecture-demo/app/user"
	"log"

	"github.com/jinzhu/gorm"
)

type mysqlRepository struct {
	Conn *gorm.DB
}

func InitMysqlRepository(Conn *gorm.DB) user.Repository {
	return &mysqlRepository{Conn}
}

func (mysqlRepo *mysqlRepository) Create(userEntity *entities.User) (int64, error) {
	if err := mysqlRepo.Conn.Save(&userEntity).Error; err != nil {
		log.Print("Saving error: ", err)
	}

	return userEntity.ID, nil
}

func (mysqlRepo *mysqlRepository) Update(condition *entities.User, userEntity *entities.User) (int64, error) {
	if err := mysqlRepo.Conn.Model(&condition).Updates(&userEntity).Error; err != nil {
		log.Print("Saving error: ", err)
	}

	return userEntity.ID, nil
}

func (mysqlRepo *mysqlRepository) GetByPhoneNumber(id int64) (entities.User, error) {
	userData := entities.User{}

	if err := mysqlRepo.Conn.First(&userData, id).Error; err != nil {
		log.Print("Saving error: ", err)
	}

	return userData, nil
}
