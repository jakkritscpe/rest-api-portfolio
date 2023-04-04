package tools

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
)

type ToolsBody struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Urlimg     string `json:"url_img"`
	CategoryID int    `json:"category_id"`
}

func ReadTools(c *gin.Context) {

	type result struct {
		ID           uint   `json:"id"`
		Name         string `json:"name"`
		Urlimg       string `json:"url_img"`
		CategoryID   int    `json:"category_id"`
		CategoryName string `json:"category_name"`
	}

	var rs []result
	db_con.Db.Table("tools").Select("tools.id, tools.name, tools.urlimg, tools.category_id, category_tools.name as category_name ").Joins("left join category_tools on category_tools.id = tools.category_id").Scan(&rs)
	c.JSON(http.StatusOK, gin.H{
		"massage": "Tools Read Success.", "data": rs,
	})
}

// If the request body is valid JSON, then check if the tool exists, if it doesn't then create it
func AddTools(c *gin.Context) {
	var json ToolsBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check user exists
	var toolsExist models.Tools
	db_con.Db.Where("name = ? ", json.Name).First(&toolsExist)
	if toolsExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Tools Exists.",
		})
		return
	}

	//Create Tools
	tools := models.Tools{Name: json.Name, Urlimg: json.Urlimg, CategoryID: json.CategoryID}
	db_con.Db.Create(&tools)
	if tools.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"massage": "Tools Created Success.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Tools Created Failed.",
		})
	}
}

// UpdateTools() is a function that receives a JSON object from the client, checks if the ID exists in
// the database, and if it exists, updates the database
func UpdateTools(c *gin.Context) {
	var json ToolsBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists
	var ToolsExist models.Tools
	db_con.Db.Where("id = ?", json.ID).First(&ToolsExist)
	if ToolsExist.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "ID Not found",
		})
		return
	}

	//Check name exists
	var toolsNameExist models.Tools
	db_con.Db.Where("name = ?", json.Name).First(&toolsNameExist)
	input_json := strings.ToUpper(toolsNameExist.Name)
	output_json := strings.ToUpper(json.Name)
	if input_json == output_json {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Name exists.",
		})
		return
	}

	//Update  Tools
	common := models.CommonFields{ID: json.ID}
	tools := models.Tools{CommonFields: common, Name: json.Name, Urlimg: json.Urlimg, CategoryID: json.CategoryID}
	err := db_con.Db.Save(&tools).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Updated Failed.",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "Updated Success.",
		})
	}
}

// I want to delete the data in the database by ID
func DeleteTools(c *gin.Context) {
	var json ToolsBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists.
	common := models.CommonFields{ID: json.ID}
	tools := models.Tools{CommonFields: common, Name: json.Name, Urlimg: json.Urlimg, CategoryID: json.CategoryID}
	err := db_con.Db.Where("id = ?", json.ID).Delete(&tools)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Delete Failed.",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"massage": "Delete success.",
		})
	}
}
