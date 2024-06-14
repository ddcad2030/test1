package routes

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Coud not parse request data. Try again later",
		})
		return
	}
	user.ID = 1
	user.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message": "User create successful",
	})
}
