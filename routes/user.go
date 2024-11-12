package routes

import (
	"net/http"

	"bookingEvent.api/models"
	"bookingEvent.api/utils"
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

func login(ctx *gin.Context) {

	// bind data
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	// validate crediantils
	err = user.ValidateCredentails()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	//  here start generate token with jwt to loged in user
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user "})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "login successfuly", "token": token})

}
