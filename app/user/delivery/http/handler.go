package http

import (
	"go-clean-architecture-demo/app/entities"
	"go-clean-architecture-demo/app/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUseCase user.UseCase
}

type CreateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type UpdateRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

func (userHandler *UserHandler) Create(c *gin.Context) {
	var requestData CreateRequest
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

	user := &entities.User{
		PhoneNumber: requestData.PhoneNumber,
		Password:    requestData.Password,
		Name:        requestData.Name,
	}

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

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data": map[string]interface{}{
			"id":          createdUser.ID,
			"name":        createdUser.Name,
			"phoneNumber": createdUser.PhoneNumber,
		},
	})
	return
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
		"data": map[string]interface{}{
			"id":          updatedUser.ID,
			"name":        updatedUser.Name,
			"phoneNumber": updatedUser.PhoneNumber,
		},
	})
	return
}

func NewEndpointHttpHandler(e *gin.Engine, us user.UseCase) {
	handler := &UserHandler{
		UserUseCase: us,
	}

	e.POST("/users", handler.Create)
	e.PUT("/users/:id", handler.Update)
}
