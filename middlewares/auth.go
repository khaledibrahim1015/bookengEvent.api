package middlewares

import (
	"net/http"

	"bookingEvent.api/utils"
	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	// verfiying token
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		// ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized ."})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized ."})
		return
	}

	userId, err := utils.VerfiyToken(token)
	if err != nil {
		// ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized ."})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized ."})
		return
	}

	//  pass value in ctx
	ctx.Set("userId", userId)

	// chain to next
	ctx.Next()

}
