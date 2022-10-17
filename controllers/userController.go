package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"

	"net/http"
	"strconv"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Age": User.Age,
		"Email":     User.Email,
		"id":        User.ID,
		"Username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email or password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	userIdParam, _ := strconv.Atoi(c.Param("userId"))


	User := models.User{}

	var newUser models.User

	if contentType == appJSON {
		c.ShouldBindJSON(&newUser)
	} else {
		c.ShouldBind(&newUser)
	}

	err := db.Model(&User).Where("id = ?", userIdParam).Take(&User).Updates(models.User{Email: newUser.Email, Username: newUser.Username}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        User.ID,
		"Email":     User.Email,
		"Username": User.Username,
		"Age": User.Age,
		"Updated_at" : User.UpdatedAt,
	})
}

func UserDelete(c *gin.Context) {
	db := database.GetDB()

	User := models.User{}
	userId, _ := strconv.Atoi(c.Param("userId"))

	err := db.Where("id = ?", userId).Delete(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}

