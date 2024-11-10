package routes

import (
	"net/http"
	"strconv"

	"bookingEvent.api/models"
	"github.com/gin-gonic/gin"
)

func getEvenets(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not fetch events . try again later !"})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	ctx.JSON(http.StatusOK, gin.H{"Data": events})

}

func getEvent(ctx *gin.Context) {
	//  get event id
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could Not parse  event id ."})
		return
	}

	//  fetch event with id
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not fetch  event ."})
		return
	}
	ctx.JSON(http.StatusCreated, event)

}

func createEvent(ctx *gin.Context) {

	var event models.Event

	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request ", "Error": err.Error()})
		return
	}

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not create event . try again later !"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created !", "Data": event})

}
