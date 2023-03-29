package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	AuthController "github.com/jakkritscpe/rest-api-portfolio/controller/auth"
	UserController "github.com/jakkritscpe/rest-api-portfolio/controller/user"
	DatabaseCon "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/middleware"

)

func main() {


	DatabaseCon.InitDB()
	log.Println("Start API portfolio ... let go !!")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", AuthController.Root)
	r.POST("/login", AuthController.Login)

	authorized := r.Group("/user", middleware.JWTAuthen())
	authorized.GET("/readall", UserController.ReadAll)
	authorized.POST("/register", AuthController.Register)

	r.Run() // listen and serve on 0.0.0.0:8080

}
