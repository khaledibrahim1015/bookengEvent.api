package main

import (
	"bookingEvent.api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.Run(":8080")

}
