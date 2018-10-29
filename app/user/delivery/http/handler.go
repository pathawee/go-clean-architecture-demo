package http

import (
	"go-clean-architecture-demo/app/user"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUseCase user.UseCase
}

func NewEndpointHttpHandler(e *gin.Engine, us user.UseCase) {
	handler := &UserHandler{
		UserUseCase: us,
	}

	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.Update)
	e.GET("/users", handler.List)
}
