package main

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture-demo/app/database"
	userHttp "go-clean-architecture-demo/app/user/delivery/http"
	userRepository "go-clean-architecture-demo/app/user/repository"
	userUseCase "go-clean-architecture-demo/app/user/usecase"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func sayPongJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", sayHello)
	r.GET("/ping", sayPongJSON)
	dbConn := database.ConnectDB()
	defer dbConn.Close()
	userRepo := userRepository.InitMysqlRepository(dbConn)
	uc := userUseCase.InitUseCase(userRepo)
	userHttp.NewEndpointHttpHandler(r, uc)
	database.DBMigration()
	r.Run()
}
