package http_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-clean-architecture-demo/app/entities"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	userHttp "go-clean-architecture-demo/app/user/delivery/http"
	"go-clean-architecture-demo/app/user/mocks"
)

func TestUserHandler_Create(t *testing.T) {
	mockUser := entities.User{
		PhoneNumber: "1234567890",
		Password:    "1234",
		Name:        "Ake",
	}

	tempMockArticle := mockUser
	tempMockArticle.ID = 0
	mockUseCase := new(mocks.UseCase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUseCase.On("Create", mock.AnythingOfType("*entities.User")).Return(&mockUser, nil)

	r := gin.Default()
	userHttp.NewEndpointHttpHandler(r, mockUseCase)

	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUseCase.AssertExpectations(t)
}
