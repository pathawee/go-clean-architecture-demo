package http

import (
	"demo/app/models"
	"demo/app/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUseCase user.UseCase
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var userEntity models.User
	err := c.Bind(&userEntity)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		}, )

		return
	}

	ar, err := userHandler.UserUseCase.Create(&userEntity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": ar,
	})
}

func NewEndpointHttpHandler(e *gin.Engine, us user.UseCase) {
	handler := &UserHandler{
		UserUseCase: us,
	}

	e.POST("/users", handler.Create)
}
