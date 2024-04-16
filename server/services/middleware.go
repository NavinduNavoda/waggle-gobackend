package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func RequireAuth(c *gin.Context) {
	fmt.Println("Authenticating user...")

	// Get the value of the "sid" parameter from JSON or form data
	sid := c.PostForm("sid") // Assuming "sid" is sent as form data
	if sid == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "SID is required"})
		c.Abort()
		return
	}

	// Perform your authentication logic here
	// For example, check if the SID is valid by querying your database or external service
	// Replace this with your actual authentication logic

	// For demonstration, let's assume SID is valid if it's not empty
	if CheckSessionId(sid) {
		// SID is valid, proceed with the next middleware or handler
		c.Next()
	} else {
		// SID is invalid, return unauthorized error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid SID"})
		c.Abort()
	}
	

}