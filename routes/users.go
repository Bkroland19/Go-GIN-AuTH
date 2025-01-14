package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)


func Signup(c *gin.Context){
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
		"message": "could not add user",
		"error":   err.Error(),
	})

	return
	}


	err = user.Save()
	if err != nil {
    c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"message": "could not save user",
		"error":   err.Error(),
	})

	return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "user created",
		"user":    user, 
	})
}


func Login(c *gin.Context){
    var user models.User


    err := c.ShouldBindJSON(&user)

    if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
		"message": "could not parse request",
		"error":   err.Error(),
	})

	return
	}

	err = user.ValidateCredentials()
	fmt.Print(err)

	if err != nil {
		fmt.Print("tru..........")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	// Generate a JWT token
	token,err := utils.GenerateToken(user.Email , user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not generate token",
			"error":   err.Error(),
			})
	}

	c.IndentedJSON(http.StatusOK , gin.H{
	  "message": "user logged in",
	  "token":   token,
	})

}