package middleware

import (
	"net/http"
	"task5-pbi-btpns-ArifBudiantoPratomo/database"
	"task5-pbi-btpns-ArifBudiantoPratomo/models"

	"github.com/gin-gonic/gin"
)

func AuthPhoto(c *gin.Context) {
	photoID := c.Param("id")
	reqID := c.GetUint("reqID")

	var photo models.Photo
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ID := photo.UserID

	if reqID != ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You dont have access"})
		return
	}

	c.Next()
}
