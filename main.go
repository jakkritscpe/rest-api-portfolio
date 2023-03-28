package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	AuthController "github.com/jakkritscpe/rest-api-portfolio/controller/auth"
	UserController "github.com/jakkritscpe/rest-api-portfolio/controller/user"
	DatabaseCon "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DatabaseCon.InitDB()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.GET("/users", UserController.ReadAllUsers)
	r.Run() // listen and serve on 0.0.0.0:8080

}
