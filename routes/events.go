package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

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
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": event})
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
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event"})
	}
	_, err = models.GetEventByID(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	var updateEvent models.Event
	err = c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update Successful", "data": updateEvent})

}
func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	element, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
	}

	err = element.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete Successful"})

}
func registerEvent(c *gin.Context) {

}
func cancelRegisterEvent(c *gin.Context) {

}

func login(c *gin.Context) {

}
