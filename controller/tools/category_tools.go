package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
)

type CategoryToolsBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// If the user exists, return an error. If the user doesn't exist, create the user
func AddCategoryTools(c *gin.Context) {
	var json CategoryToolsBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check user exists
	var categoryToolsExist models.CategoryTools
	db_con.Db.Where("name = ?", json.Name).First(&categoryToolsExist)
	if categoryToolsExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Category Tools Exists.",
		})
		return
	}

	//Create Category Tools
	categoryTools := models.CategoryTools{Name: json.Name}
	db_con.Db.Create(&categoryTools)
	if categoryTools.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"massage": "Category Tools Created Success.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Category Tools Created Failed.",
		})
	}
}

// UpdateCategoryTools() is a function that is used to update the data in the database
func UpdateCategoryTools(c *gin.Context) {
	var json CategoryToolsBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists
	var categoryToolsExist models.CategoryTools
	db_con.Db.Where("id = ?", json.ID).First(&categoryToolsExist)
	if categoryToolsExist.ID < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "ID Not found",
		})
		return
	}

	//Create Category Tools
	categoryTools := models.CategoryTools{ID: json.ID, Name: json.Name}
	err := db_con.Db.Save(&categoryTools).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Updated Failed.",
		})

	} else {
		c.JSON(http.StatusCreated, gin.H{
			"massage": "Updated Success.",
		})
	}
}

// I want to delete a record from the database by ID
func DeleteCategoryTools(c *gin.Context) {
	var json CategoryToolsBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists
	categoryTools := models.CategoryTools{ID: json.ID, Name: json.Name}
	err := db_con.Db.Where("id = ?", json.ID).Delete(&categoryTools)
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
