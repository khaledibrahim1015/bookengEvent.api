package routes

import (
	"net/http"
	"strconv"

	"bookingEvent.api/models"
	"bookingEvent.api/utils"
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

	// verfiying token
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized ."})
		return
	}

	userId, err := utils.VerfiyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized ."})
		return
	}

	var event models.Event
	err = ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request ", "Error": err.Error()})
		return
	}

	event.UserId = userId
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not create event . try again later !"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "event created !", "Data": event})

}

// update
func updateEvent(ctx *gin.Context) {

	//  get this event
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could Not parse  event id ."})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could Not fetch the event."})
		return
	}

	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request ", "Error": err.Error()})
		return
	}

	//  update event
	// assign id
	updatedEvent.ID = eventId
	err = updatedEvent.Updated()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not update event . try again later !"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfuly !"})
}

// delete
func deleteEvent(ctx *gin.Context) {

	//  get current id
	eventid, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could Not parse  event id ."})
		return
	}

	//  get current event
	event, err := models.GetEventById(eventid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not fetch the event."})
		return
	}

	//  delete event
	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Delete The Event ."})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfuly !"})

}
