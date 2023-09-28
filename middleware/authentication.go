package middleware

import (
	"net/http"
	"strings"
	"task5-pbi-btpns-ArifBudiantoPratomo/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token missing"})
		return
	}

	accessToken := strings.Split(authHeader, " ")[1]
	claims, err := helpers.ReadToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set("reqID", claims.ID)

	c.Next()
}
