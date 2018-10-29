package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	Skip  int `json:"skip" validate:"required"`
	Limit int `json:"limit" validate:"required"`
}

func (userHandler *UserHandler) List(c *gin.Context) {
	var requestData ListRequest
	err := c.Bind(&requestData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

		return
	}

	users, err := userHandler.UserUseCase.List(requestData.Skip, requestData.Limit)
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
		"data":   users,
	})
	return
}
