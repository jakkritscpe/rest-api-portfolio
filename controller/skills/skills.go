package skills

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
)

type SkillBody struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ReadSkills(c *gin.Context) {
	var skills []models.Skills
	db_con.Db.Find(&skills)
	c.JSON(http.StatusOK, gin.H{
		"massage": "Skills Read Success.", "data": skills,
	})
}

// If the request body is valid JSON, check if the skill already exists, if it doesn't, create it
func AddSkill(c *gin.Context) {
	var json SkillBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check Skills exists
	var skillExist models.Skills
	db_con.Db.Where("name = ?", json.Name).First(&skillExist)
	if skillExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Skills Exists.",
		})
		return
	}

	//Create Skills
	Skill := models.Skills{Name: json.Name}
	db_con.Db.Create(&Skill)
	if Skill.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"massage": "Skill Created Success.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Skill Created Failed.",
		})
	}
}

// UpdateSkill() is a function that takes a JSON object from the client, checks if the ID exists in the
// database, and if it does, updates the record
func UpdateSkill(c *gin.Context) {
	var json SkillBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check exists
	var skillExist models.Skills
	db_con.Db.Where("id = ?", json.ID).First(&skillExist)
	if skillExist.ID < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "ID Not found",
		})
		return
	}

	//Check name exists
	var skillNameExist models.Skills
	db_con.Db.Where("name = ?", json.Name).First(&skillNameExist)
	input_json := strings.ToUpper(skillNameExist.Name)
	output_json := strings.ToUpper(json.Name)
	if input_json == output_json {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Name exists.",
		})
		return
	}

	//Update Skills
	Skill := models.Skills{ID: json.ID, Name: json.Name}
	err := db_con.Db.Save(&Skill).Error
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
func DeleteSkill(c *gin.Context) {
	var json SkillBody

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Delete Skills
	skill := models.Skills{ID: json.ID, Name: json.Name}
	err := db_con.Db.Where("id = ?", json.ID).Delete(&skill)
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
