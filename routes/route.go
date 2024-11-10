package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	fmt.Println("Registering routes...") // Debug print
	server.GET("/events", getEvenets)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	fmt.Println("Routes registered.") // Debug print
}
