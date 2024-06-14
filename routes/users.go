package routes

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"

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
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Coud not save user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User create successful",
	})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data. Try again later",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user 1."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}
