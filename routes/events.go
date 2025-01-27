package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)



func GetEvent(c*gin.Context){
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

func GetEvents(c *gin.Context){
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not get events",
	})
		return

	}
	c.JSON(http.StatusOK, events)

	
}



func CreateEvent(c *gin.Context){
	
	var event models.Event
	err := c.ShouldBindJSON(&event)
    
	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
		"message": "could not add event",
		"error":   err.Error(),
	})

	return
	}

	userId := c.GetInt64("userId")
	// fmt.Print(userId,"===")
	event.UserId = userId

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

func UpdateEvent(c *gin.Context){
	eventId, err := strconv.ParseInt((c.Param("id")),10,64)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event id",
		})

		return
	}

	userId := c.GetInt64("userId")
	event , err := models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid event id",
		})

		return
	}

	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	var UpdatedEvent models.Event

	err = c.ShouldBindJSON(&UpdatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event data",
		})

		return
	}

	UpdatedEvent.ID = eventId


	err = UpdatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not update event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "event updated",
		"event": UpdatedEvent,
	})

}

func DeleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	userId := c.GetInt64("userId")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event id",
		})

		return
	}

	var event *models.Event
	event, err = models.GetEventById(eventId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "not such event",
		})

		return
	}


	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not delete event",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "event deleted",
	})

}


