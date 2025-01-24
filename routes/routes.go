package routes

import (
	"github.com/gin-gonic/gin"

	"example.com/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine){

	server.GET("/events", GetEvents )
	server.GET("/events/:id", GetEvent)
	
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreateEvent)
    authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent) 
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)


	//user routes
	server.POST("/signup",Signup)
	server.POST("/login",Login)
}