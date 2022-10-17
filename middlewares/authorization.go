package middlewares

import (
	"net/http"
	"mygram/database"
	"mygram/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func (c *gin.Context){
		db := database.GetDB()
		userIdParam, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
				"message": "invalid param",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err = db.Select("id").First(&User, uint(userIdParam)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Not found",
				"message" : "data not exist",
			})
			return
		}

		if User.ID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
				"message" : "you are not authorized to access this data",
			})
			return
		}
		c.Next()	
	}
}

// func ProductAuthorization() gin.HandlerFunc {
// 	return func (c *gin.Context){
// 		db := database.GetDB()
// 		productID, err := strconv.Atoi(c.Param("productId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error": "Bad request",
// 				"message": "invalid param",
// 			})
// 			return
// 		}
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		userID := uint(userData["id"].(float64))
// 		Product := models.Product{}

// 		err = db.Select("user_id").First(&Product, uint(productID)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error": "Not found",
// 				"message" : "data not exist",
// 			})
// 			return
// 		}

// 		if Product.UserID != userID {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error" : "Unauthorized",
// 				"message" : "you are not authorized to access this data",
// 			})
// 			return
// 		}
// 		c.Next()	
// 	}
// }