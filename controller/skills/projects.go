package skills

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
)

type ProjectBody struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Urlimg      string `json:"url_img"`
	Urlvideo    string `json:"url_video"`
	SkillID     int    `json:"skill_id"`
}

func ReadProjects(c *gin.Context) {

	type result struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Urlimg      string `json:"url_img"`
		Urlvideo    string `json:"url_video"`
		SkillID     int    `json:"skill_id"`
		SkillName   string `json:"skill_name"`
	}

	var rs []result
	db_con.Db.Table("projects").Select("projects.id, projects.name, projects.urlimg , projects.urlvideo, projects.skill_id, skills.name as skill_name ").Joins("left join skills on skills.id = projects.skill_id").Scan(&rs)
	c.JSON(http.StatusOK, gin.H{
		"massage": "Projects Read Success.", "data": rs,
	})
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
			"massage": "Project name Exists.",
		})
		return
	}

	//Create Projects
	project := models.Projects{Name: json.Name, Urlimg: json.Urlimg, Urlvideo: json.Urlvideo, SkillID: json.SkillID}
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

	//Check name exists
	var projectNameExist models.Projects
	db_con.Db.Where("name = ?", json.Name).First(&projectNameExist)
	input_json := strings.ToUpper(projectNameExist.Name)
	output_json := strings.ToUpper(json.Name)
	if input_json == output_json {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Name exists.",
		})
		return
	}

	//Update  Projects
	common := models.CommonFields{ID: json.ID}
	project := models.Projects{CommonFields: common, Name: json.Name, Urlimg: json.Urlimg, Urlvideo: json.Urlvideo, SkillID: json.SkillID}
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
	project := models.Projects{CommonFields: common, Name: json.Name, Urlimg: json.Urlimg, Urlvideo: json.Urlvideo, SkillID: json.SkillID}
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
