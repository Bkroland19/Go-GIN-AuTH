package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/models"
)



func registerForEvent(c*gin.Context){
	userId := c.GetInt64("userId");

	eventId, err := strconv.ParseInt((c.Param("id")),10,64)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event id",
		})

		return
	}



	 event, err := models.GetEventById(eventId);

	 if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"internal server error",
		})
	 }


	err = event.Register(userId)


	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"message":"could not register user for the event",
		})

		return
	}


	c.IndentedJSON(http.StatusCreated,gin.H{
		"message":"user registered for event",
	})


}

func cancelRegistration(c*gin.Context){
	userId := c.GetInt64("userId");
	eventId, err := strconv.ParseInt((c.Param("id")),10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid event id",
			})
		} 

		var event models.Event
		event.ID = eventId

		err = event.CancelRegistration(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not cancel registration",
				})
				return
				}
				c.IndentedJSON(http.StatusOK, gin.H{
					"message": "registration cancelled",
					})	
}