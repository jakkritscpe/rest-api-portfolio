package skills

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
)

type ProjectBody struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Urlimg      string `json:"url_img"`
	SkillID     int    `json:"category_id"`
}

func AddProject(c *gin.Context) {
	var json ProjectBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check user exists
	var projectExist models.Projects
	db_con.Db.Where("name = ? ", json.Name).First(&projectExist)
	if projectExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Tools Exists.",
		})
		return
	}

	//Create Projects
	project := models.Projects{Name: json.Name, Urlimg: json.Urlimg, SkillID: json.SkillID}
	db_con.Db.Create(&project)
	if project.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"massage": "Project Created Success.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Project Created Failed.",
		})
	}
}

func UpdateProject(c *gin.Context) {
	var json ProjectBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists
	var projectExist models.Projects
	db_con.Db.Where("id = ?", json.ID).First(&projectExist)
	if projectExist.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "ID Not found",
		})
		return
	}

	//Update  Projects
	common := models.CommonFields{ID: json.ID}
	project := models.Projects{CommonFields: common, Name: json.Name, Urlimg: json.Urlimg, SkillID: json.SkillID}
	err := db_con.Db.Save(&project).Error
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

func DeleteProject(c *gin.Context) {
	var json ProjectBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists.
	common := models.CommonFields{ID: json.ID}
	project := models.Projects{CommonFields: common, Name: json.Name, Urlimg: json.Urlimg, SkillID: json.SkillID}
	err := db_con.Db.Where("id = ?", json.ID).Delete(&project)
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
