package routes

import (
	"net/http"

	"example.com/go-rest/constants"
	"example.com/go-rest/models"
	"example.com/go-rest/utils"
	"github.com/gin-gonic/gin"
)

const (
	msgRegisterEvent      = "Could not register event."
	msgCancelRegistration = "Could not cancel event registration."
	msgRegisteredOK       = "Registered Event!"
	msgCancelledOK        = "Cancelled Registration!"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64(constants.ContextUserID)
	eventId, err := utils.ParseInt64(context.Param(constants.ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseEventID})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgFetchEvent})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgRegisterEvent})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": msgRegisterEvent})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64(constants.ContextUserID)
	eventId, err := utils.ParseInt64(context.Param(constants.ID))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseEventID})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgParseEventID})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": msgCancelledOK})

}
