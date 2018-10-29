package http

import (
	"go-clean-architecture-demo/app/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var requestData CreateRequest

	// Request validation & Map params to request data
	err := c.ShouldBind(&requestData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"data": map[string]interface{}{
				"message": err.Error(),
			},
		})

		return
	}

	// Map request data to entity (sanitize request)
	user := &entities.User{
		PhoneNumber: requestData.PhoneNumber,
		Password:    requestData.Password,
		Name:        requestData.Name,
	}

	// Create user
	createdUser, err := userHandler.UserUseCase.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"data": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"id":          createdUser.ID,
			"name":        createdUser.Name,
			"phoneNumber": createdUser.PhoneNumber,
		},
	})
	return
}
