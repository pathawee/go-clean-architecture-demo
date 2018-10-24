package main

import (
	"demo/app/database"
	userHttp "demo/app/user/delivery/http"
	userRepository "demo/app/user/repository"
	userUseCase "demo/app/user/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func sayHello(c *gin.Context) {
	log.Print("env POSTGRES_DB: ", os.Getenv("POSTGRES_DB"))
	c.String(http.StatusOK, "Hello " + os.Getenv("POSTGRES_DB"))
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

	//// get port env var
	//port := "8080"
	//portEnv := os.Getenv("PORT")
	//if len(portEnv) > 0 {
	//	port = portEnv
	//}
	//
	//log.Printf("Listening on port %s...", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	//
}
