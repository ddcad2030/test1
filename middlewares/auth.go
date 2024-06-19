package middlewares

import (
	"net/http"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
	}

	c.Set("userId", userId)
	c.Next()
}
