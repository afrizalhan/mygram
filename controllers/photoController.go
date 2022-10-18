package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"strconv"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreatePhoto godoc
// @Summary Create new photo data.
// @Description User post a picture.
// @Tags Photos
// @Param Body body models.PhotoInput true "the body to create a photo"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /photos/ [post]
func CreatePhoto(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	err := db.Debug().Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
		})
		return
	}

	Photo.User = User

	err = db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id" : Photo.ID,
		"title": Photo.Title,
		"caption": Photo.Caption,
		"photo_url": Photo.PhotoURL,
		"user_id": Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

// GetPhotos godoc
// @Summary Get all Photo from database.
// @Description Get every photo that exist on MyGram.
// @Tags Photos
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.GetPhotoRes
// @Router /photos [get]
func GetPhotos(c *gin.Context) {
	db := database.GetDB()

	photos := []models.Photo{}

	err := db.Preload("User").Find(&photos).Error

	result := []models.GetPhotoRes{}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	for _, photo := range photos {
		res := models.GetPhotoRes{}

		res.ID = photo.ID
		res.Title = photo.Title
		res.Caption = photo.Caption
		res.PhotoURL = photo.PhotoURL
		res.UserID = photo.UserID
		res.CreatedAt = photo.CreatedAt
		res.UpdatedAt = photo.UpdatedAt
		res.User.Email = photo.User.Email
		res.User.Username = photo.User.Username

		result = append(result, res)
	}

	c.JSON(http.StatusOK, result)
}

// UpdatePhoto godoc
// @Summary Update Photo.
// @Description Update Photo by id.
// @Tags Photos
// @Produce json
// @Param photoId path string true "Photo id"
// @Param Body body models.PhotoInput true "the body to update photo"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]interface{}
// @Router /photos/{photoId} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	User := models.User{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Debug().Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Unauthorized",
		})
		return
	}

	Photo.User = User

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id" : Photo.ID,
		"title": Photo.Title,
		"caption": Photo.Caption,
		"photo_url": Photo.PhotoURL,
		"user_id": Photo.UserID,
		"updated_at": Photo.UpdatedAt,
	})
}

// DeletePhoto godoc
// @Summary Delete one Photo .
// @Description Delete a Photo by id.
// @Tags Photos
// @Produce json
// @Param photoId path string true "Photo id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /photos/{photoId} [delete]
func DeletePhoto(c *gin.Context) {
	db := database.GetDB()

	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	err := db.Where("id = ?", photoId).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
