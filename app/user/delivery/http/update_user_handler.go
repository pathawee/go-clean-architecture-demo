package http

import (
	"go-clean-architecture-demo/app/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

func (userHandler *UserHandler) Update(c *gin.Context) {
	var requestData UpdateRequest
	err := c.Bind(&requestData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	user := &entities.User{
		PhoneNumber: requestData.PhoneNumber,
		Name:        requestData.Name,
	}
	updatedUser, err := userHandler.UserUseCase.UpdateById(int64(userID), user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"data": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": gin.H{
			"id":          updatedUser.ID,
			"name":        updatedUser.Name,
			"phoneNumber": updatedUser.PhoneNumber,
		},
	})
	return
}
