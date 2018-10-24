package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-clean-architecture-demo/app/entities"
	"go-clean-architecture-demo/app/user/mocks"
	"go-clean-architecture-demo/app/user/usecase"
	"testing"
)

func TestUseCase_Create(t *testing.T) {
	mockUserRepo := new(mocks.Repository)
	mockUser := entities.User{
		ID:          99,
		PhoneNumber: "1234567890",
		Password:    "1234",
		Name:        "Ake",
	}

	mockUserRepo.On("Create", mock.AnythingOfType("*entities.User")).Return(mockUser.ID, nil)

	u := usecase.InitUseCase(mockUserRepo)

	a, err := u.Create(&mockUser)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	mockUserRepo.AssertExpectations(t)
}
