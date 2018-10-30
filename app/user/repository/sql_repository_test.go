package repository_test

import (
	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-demo/app/entities"
	"go-clean-architecture-demo/app/user/repository"
	"log"
	"testing"
)

func SetupDBTests() *gorm.DB {
	mocket.Catcher.Register()

	db, err := gorm.Open(mocket.DRIVER_NAME, "")
	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}
	// Log mode shows the query gorm uses, so we can replicate and mock it
	db.LogMode(true)

	return db
}

func TestMysqlRepository_Create(t *testing.T) {
	userEntity := &entities.User{
		PhoneNumber: "1234567890",
		Password:    "1234",
		Name:        "Ake",
	}

	db := SetupDBTests()
	defer db.Close()

	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"users\"").WithId(3)
	repo := repository.InitMysqlRepository(db)

	lastId, err := repo.Create(userEntity)
	assert.NoError(t, err)
	assert.Equal(t, uint(3), lastId)
}
