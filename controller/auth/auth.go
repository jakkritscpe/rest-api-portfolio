package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	db_con "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte

// Service help check
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"massge": "Hi this is API portfolio."})
}


// Service Register
type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check user exists
	var userExist models.User
	db_con.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "error", "massage": "User Exists.",
		})
		return
	}

	//Create User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := models.User{Username: json.Username, Password: string(encryptedPassword), Fullname: json.Fullname, Nickname: json.Nickname, Avatar: json.Avatar}
	db_con.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"status": "ok", "massage": "User Created Success.", "userID": user.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "error", "massage": "User Created Failed.",
		})
	}
}

// Service Login
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check User Exists
	var userExist models.User
	db_con.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "error", "massage": "User Dose Not Exists.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 10).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		log.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{
			"status": "ok", "massage": "Login Success.", "token": tokenString,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "error", "massage": "Login Failed.",
		})
	}

}
