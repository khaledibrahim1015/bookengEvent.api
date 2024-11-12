package routes

import (
	"bookingEvent.api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvenets)
	server.GET("/events/:id", getEvent)

	authenticatedRouterGroup := server.Group("/")
	authenticatedRouterGroup.Use(middlewares.Authentication)
	authenticatedRouterGroup.POST("/events", createEvent)
	authenticatedRouterGroup.PUT("/events/:id", updateEvent)
	authenticatedRouterGroup.DELETE("/events/:id", deleteEvent)

	// server.POST("/events", middlewares.Authentication, createEvent)
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
