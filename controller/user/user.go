package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
)

func ReadAll(c *gin.Context) {
	var users []models.User
	db_con.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok", "massage": "User Read Success.", "users": users,
	})
}
