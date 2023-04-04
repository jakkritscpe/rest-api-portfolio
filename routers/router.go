package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jakkritscpe/rest-api-portfolio/middleware"

	AuthController "github.com/jakkritscpe/rest-api-portfolio/controller/auth"
	SkillsController "github.com/jakkritscpe/rest-api-portfolio/controller/skills"
	ToolsController "github.com/jakkritscpe/rest-api-portfolio/controller/tools"
	UserController "github.com/jakkritscpe/rest-api-portfolio/controller/user"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", AuthController.Root)
	r.POST("/login", AuthController.Login)

	r.GET("/categorytools", ToolsController.ReadCategoryTools)
	r.GET("/tools", ToolsController.ReadTools)

	r.GET("/skills", SkillsController.ReadSkills)
	r.GET("/projects", SkillsController.ReadProjects)

	authorized := r.Group("/user", middleware.JWTAuthen())
	// user
	authorized.POST("/register", AuthController.Register)
	authorized.GET("/readall", UserController.ReadAll)

	// tools create
	authorized.POST("/categorytool", ToolsController.AddCategoryTools)
	authorized.POST("/tool", ToolsController.AddTools)
	// - update
	authorized.PATCH("/categorytool", ToolsController.UpdateCategoryTools)
	authorized.PATCH("/tool", ToolsController.UpdateTools)
	// - delete
	authorized.DELETE("/categorytool", ToolsController.DeleteCategoryTools)
	authorized.DELETE("/tool", ToolsController.DeleteTools)

	//skills
	authorized.POST("/skill", SkillsController.AddSkill)
	authorized.POST("/project", SkillsController.AddProject)
	// - update
	authorized.PATCH("/skill", SkillsController.UpdateSkill)
	authorized.PATCH("/project", SkillsController.UpdateProject)
	// - delete
	authorized.DELETE("/skill", SkillsController.DeleteSkill)
	authorized.DELETE("/project", SkillsController.DeleteProject)

	return r
}
