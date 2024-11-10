package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvenets)
	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

}
