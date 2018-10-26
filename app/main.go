package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func sayHello(c *gin.Context) {
	c.String(http.StatusOK, "Hello "+os.Getenv("APP_ENV"))
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
	//dbConn := database.ConnectDB()
	//defer dbConn.Close()
	//userRepo := userRepository.InitMysqlRepository(dbConn)
	//uc := userUseCase.InitUseCase(userRepo)
	//userHttp.NewEndpointHttpHandler(r, uc)
	//database.DBMigration()
	r.Run()
}
