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

// CreateComment godoc
// @Summary Create new comment for a photo.
// @Description User post a comment on a picture.
// @Tags Comments
// @Param Body body models.CommentInput true "the body to create a photo"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /comments/ [post]
func CreateComment(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	Photo := models.Photo{}
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	err := db.Debug().Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	Comment.User = User

	err = db.Debug().Where("id = ?", Comment.PhotoID).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	Comment.Photo = Photo
	Comment.Photo.User = User

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
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
func GetComments(c *gin.Context) {
	db := database.GetDB()

	comments := []models.Comment{}

	err := db.Preload("User").Preload("Photo").Find(&comments).Error

	result := []models.GetCommentRes{}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	for _, comment := range comments {
		res := models.GetCommentRes{}

		res.ID = comment.ID
		res.Message = comment.Message
		res.PhotoID = comment.PhotoID
		res.UserID = comment.UserID
		res.CreatedAt = comment.CreatedAt
		res.UpdatedAt = comment.UpdatedAt
		res.User.ID = comment.User.ID
		res.User.Email = comment.User.Email
		res.User.Username = comment.User.Username
		res.Photo.ID = comment.Photo.ID
		res.Photo.Title = comment.Photo.Title
		res.Photo.Caption = comment.Photo.Caption
		res.Photo.PhotoURL = comment.Photo.PhotoURL
		res.Photo.UserID = comment.Photo.UserID

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
func UpdateComment(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	newComment := models.Comment{}
	Photo := models.Photo{}
	User := models.User{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&newComment)
	} else {
		c.ShouldBind(&newComment)
	}

	err := db.Debug().Where("id = ?", commentId).Take(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "comment not found",
		})
		return
	}

	newComment.UserID = userID
	newComment.ID = uint(commentId)

	err = db.Debug().Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	Comment.User = User

	err = db.Debug().Where("id = ?", Comment.PhotoID).Take(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	err = db.Model(&newComment).Where("id = ?", commentId).Updates(models.Comment{Message: newComment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         newComment.ID,
		"message":    newComment.Message,
		"photo_id":   Photo.ID,
		"user_id":    userID,
		"updated_at": newComment.UpdatedAt,
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
func DeleteComment(c *gin.Context) {
	db := database.GetDB()

	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.Where("id = ?", commentId).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
