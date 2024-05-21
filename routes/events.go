package routes

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	element, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event. Try again later"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": element})
}

func getEvent(c *gin.Context) {

}

func createEvent(c *gin.Context) {
	var e models.Event

	err := c.ShouldBindJSON(&e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Coud not create event. Try again later",
		})
		return
	}
	e.ID = 1
	e.UserID = 1

	e.Save()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created !", "event": e,
	})
}

func updateEvent(c *gin.Context) {

}
func deleteEvent(c *gin.Context) {

}
func registerEvent(c *gin.Context) {

}
func cancelRegisterEvent(c *gin.Context) {

}

func signup(c *gin.Context) {

}

func login(c *gin.Context) {

}
