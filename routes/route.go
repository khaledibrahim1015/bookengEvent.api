package routes

import (
	"bookingEvent.api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvenets)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authentication, createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
