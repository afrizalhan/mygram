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

// CreateSocialMedia godoc
// @Summary Create new social media for user.
// @Description User add their social media.
// @Tags Social Media
// @Param Body body models.SocialMediaInput true "the body to create a social media"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /socialmedias/ [post]
func CreateSocial(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Social := models.SocialMedia{}
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.UserID = userID
	err := db.Debug().Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "User not found",
		})
		return
	}

	Social.User = User

	err = db.Debug().Create(&Social).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id" : Social.ID,
		"name": Social.Name,
		"social_media_url": Social.SocialMediaUrl,
		"user_id": Social.UserID,
		"created_at": Social.CreatedAt,
	})
}

// GetSocials godoc
// @Summary Get all socials.
// @Description Get every socials that every user added.
// @Tags Social Media
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.GetSocialRes
// @Router /socialmedias [get]
func GetSocials(c *gin.Context) {
	db := database.GetDB()

	Socials := []models.SocialMedia{}

	err := db.Preload("User").Find(&Socials).Error

	result := []models.GetSocialRes{}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	for _, social := range Socials {
		res := models.GetSocialRes{}

		res.ID = social.ID
		res.Name = social.Name
		res.SocialMediaUrl = social.SocialMediaUrl
		res.UserID = social.UserID
		res.CreatedAt = social.CreatedAt
		res.UpdatedAt = social.UpdatedAt
		res.User.ID = social.User.ID
		res.User.Email = social.User.Email
		res.User.Username = social.User.Username

		result = append(result, res)
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": result,
	})
}

// UpdateSocial godoc
// @Summary Update Social Media.
// @Description Update added social media by id.
// @Tags Social Media
// @Produce json
// @Param socialMediaId path string true "Social Media id"
// @Param Body body models.SocialMediaInput true "the body to update social media"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]interface{}
// @Router /socialmedias/{socialMediaId} [put]
func UpdateSocial(c *gin.Context) {
	db := database.GetDB()

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Social := models.SocialMedia{}
	User := models.User{}
	socialId, _ := strconv.Atoi(c.Param("socialMediaId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Social)
	} else {
		c.ShouldBind(&Social)
	}

	Social.UserID = userID
	Social.ID = uint(socialId)

	err := db.Debug().Where("id = ?", userID).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "User not found",
		})
		return
	}

	Social.User = User

	err = db.Model(&Social).Where("id = ?", socialId).Updates(models.SocialMedia{Name: Social.Name, SocialMediaUrl: Social.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id" : Social.ID,
		"name": Social.Name,
		"social_media_url": Social.SocialMediaUrl,
		"user_id": Social.UserID,
		"updated_at": Social.UpdatedAt,
	})
}

// DeleteSocial godoc
// @Summary Delete added social media .
// @Description Delete a Social media by id.
// @Tags Social Media
// @Produce json
// @Param socialMediaId path string true "Social Media id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /socialmedias/{socialMediaId} [delete]
func DeleteSocial(c *gin.Context) {
	db := database.GetDB()

	Social := models.SocialMedia{}
	socialId, _ := strconv.Atoi(c.Param("socialMediaId"))

	err := db.Where("id = ?", socialId).Delete(&Social).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
