package routes

import (
	"net/http"

	"bookingEvent.api/models"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {

	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user !"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": " user created successfuly"})
}