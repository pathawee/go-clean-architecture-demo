package http

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture-demo/app/entities"
	"go-clean-architecture-demo/app/user"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUseCase user.UseCase
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var userEntity entities.User
	err := c.Bind(&userEntity)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})

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
