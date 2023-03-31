package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jakkritscpe/rest-api-portfolio/middleware"

	AuthController "github.com/jakkritscpe/rest-api-portfolio/controller/auth"
	ToolsController "github.com/jakkritscpe/rest-api-portfolio/controller/tools"
	UserController "github.com/jakkritscpe/rest-api-portfolio/controller/user"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", AuthController.Root)
	r.POST("/login", AuthController.Login)

	authorized := r.Group("/user", middleware.JWTAuthen())

	// user
	authorized.POST("/register", AuthController.Register)
	authorized.GET("/readall", UserController.ReadAll)

	// tools create
	authorized.POST("/categorytools", ToolsController.AddCategoryTools)
	authorized.POST("/tools", ToolsController.AddTools)
	// - update
	authorized.PATCH("/categorytools", ToolsController.UpdateCategoryTools)
	authorized.PATCH("/tools", ToolsController.UpdateTools)
	// - delete
	authorized.DELETE("/categorytools", ToolsController.DeleteCategoryTools)
	authorized.DELETE("/tools", ToolsController.DeleteTools)

	//skills

	return r
}
