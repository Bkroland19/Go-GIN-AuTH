package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){

	server.GET("/events", GetEvents )
	server.GET("/events/:id", GetEvent)
	server.POST("/events",middlewares.Authenticate, CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent) 

	//user routes
	server.POST("/signup",Signup)
	server.POST("/login",Login)
}