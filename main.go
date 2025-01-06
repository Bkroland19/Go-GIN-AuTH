package main

import (
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)




func main(){

	db.InitDB()
	server := gin.Default()


	server.GET("/events", getEvents )
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)


	server.Run(":8080")
}


func getEvent(c*gin.Context){
	eventId,err := strconv.ParseInt(c.Param("id"),10,64) 

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"message": "invalid event id",
		})

		return
	}

	event,err := models.GetEventById(eventId)


	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message": "could not get event",
		})
		return
	}
	c.JSON(http.StatusOK, event)
	
}

func getEvents(c *gin.Context){
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not get events",
	})
		return

	}
	c.JSON(http.StatusOK, events)

	
}



func createEvent(c *gin.Context){
	var event models.Event
	err := c.ShouldBindJSON(&event)
    
	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
		"message": "could not add event",
		"error":   err.Error(),
	})

	return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not add event",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "event created",
		"event":event,
	})
}