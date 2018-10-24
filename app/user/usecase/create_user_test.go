package usecase_test

import (
	"demo/app/models"
	"demo/app/user/mocks"
	"demo/app/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCase_Create(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	mockUser := models.User{
		ID: 99,
		PhoneNumber: "1234567890",
		Password: "1234",
		Name: "Ake",
	}

	mockUserRepo.On("Create", mock.AnythingOfType("*models.User")).Return(mockUser.ID, nil)

	u := usecase.InitUseCase(mockUserRepo)

	a, err := u.Create(&mockUser)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	mockUserRepo.AssertExpectations(t)
}
