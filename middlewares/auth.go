package middlewares

import (
    "net/http"


    "example.com/rest-api/utils"
    "github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "no token provided"})
        c.Abort()
        return
    }

    // Extract userId using VerifyToken
    userId, err := utils.VerifyToken(tokenString)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "error": err.Error()})
        c.Abort()
        return
    }

    // Set userId in context
    c.Set("userId", userId)
    c.Next()
}
