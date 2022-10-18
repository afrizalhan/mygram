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

// RegisterUser godoc
// @Summary Register as as user.
// @Description User Register to get access to MyGram.
// @Tags User
// @Param Body body models.RegisterInput true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/register [post]
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
		"Age":      User.Age,
		"Email":    User.Email,
		"id":       User.ID,
		"Username": User.Username,
	})
}

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access MyGram.
// @Tags User
// @Param Body body models.LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/login [post]
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

// UpdateUser godoc
// @Summary Update User.
// @Description Update User by id.
// @Tags User
// @Produce json
// @Param userId path string true "User id"
// @Param Body body models.UpdateInput true "the body to update user"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]interface{}
// @Router /users/{userId} [put]
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
		"id":         User.ID,
		"Email":      User.Email,
		"Username":   User.Username,
		"Age":        User.Age,
		"Updated_at": User.UpdatedAt,
	})
}

// DeleteUser godoc
// @Summary Delete one User .
// @Description Delete a User by id.
// @Tags User
// @Produce json
// @Param userId path string true "User id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /users/{userId} [delete]
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
