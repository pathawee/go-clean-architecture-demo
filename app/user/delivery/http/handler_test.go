package http_test

import (
	"demo/app/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	userHttp "demo/app/user/delivery/http"
	"demo/app/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_Create(t *testing.T) {
	mockUser := models.User{
		PhoneNumber: "1234567890",
		Password:    "1234",
		Name:        "Ake",
	}

	tempMockArticle := mockUser
	tempMockArticle.ID = 0
	mockUseCase := new(mocks.UseCase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUseCase.On("Create", mock.AnythingOfType("*models.User")).Return(&mockUser, nil)

	r := gin.Default()
	userHttp.NewEndpointHttpHandler(r, mockUseCase)

	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUseCase.AssertExpectations(t)
}
