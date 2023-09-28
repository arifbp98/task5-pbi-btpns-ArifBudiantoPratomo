package middleware

import (
	"net/http"
	"task5-pbi-btpns-ArifBudiantoPratomo/database"
	"task5-pbi-btpns-ArifBudiantoPratomo/models"

	"github.com/gin-gonic/gin"
)

func AuthUser(c *gin.Context) {
	userID := c.Param("id")
	reqID := c.GetUint("reqID")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ID := user.ID

	if reqID != ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You dont have access"})
		return
	}

	c.Next()
}
