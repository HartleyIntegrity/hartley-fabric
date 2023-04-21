package routes

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joe/hartley-backend/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidSigningMethod
			}
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("email", claims.Email)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}
	}
}

func CreateContract(c *gin.Context) {
	var contract models.Contract

	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	models.DB.Create(&contract)

	c.JSON(http.StatusOK, gin.H{"message": "Contract created"})
}

func GetContracts(c *gin.Context) {
	var contracts []models.Contract

	models.DB.Find(&contracts)

	c.JSON(http.StatusOK, gin.H{"contracts": contracts})
}

func UpdateContract(c *gin.Context) {
	id := c.Param("id")

	var contract models.Contract

	if err := models.DB.Where("id = ?", id).First(&contract).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	if err := c.ShouldBindJSON(&contract); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	models.DB.Save(&contract)

	c.JSON(http.StatusOK, gin.H{"message": "Contract updated"})
}

func DeleteContract(c *gin.Context) {
	id := c.Param("id")

	var contract models.Contract

	if err := models.DB.Where("id = ?", id).First(&contract).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contract not found"})
		return
	}

	models.DB.Delete(&contract)

	c.JSON(http.StatusOK, gin.H{"message": "Contract deleted"})
}
