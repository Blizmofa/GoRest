package routes

import (
	"net/http"
	"strconv"

	"example.com/go-rest/models"
	"github.com/gin-gonic/gin"
)

const (
	msgFetchEvents      = "Could not fetch events for database."
	msgParseEventID     = "Could not parse event ID."
	msgFetchEvent       = "Could not fetch event."
	msgParseRequestData = "Could not parse request data."
	msgCreateEvent      = "Could not create event."
	msgNotAuthUpdate    = "Not authorized to update event."
	msgUpdateEvent      = "Could not update event."
	msgNotAuthDelete    = "Not authorized to delete event."
	msgDeleteEvent      = "Could not delete event."

	msgEventCreated = "Event created!"
	msgEventUpdated = "Event updated!"
	msgEventDeleted = "Event deleted!"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgFetchEvents})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseEventID})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgFetchEvent})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseRequestData})
		return
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgCreateEvent})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": msgEventCreated, "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseEventID})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgFetchEvent})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": msgNotAuthUpdate})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseRequestData})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgUpdateEvent})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": msgEventUpdated, "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": msgParseEventID})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgFetchEvent})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": msgNotAuthDelete})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": msgDeleteEvent})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": msgEventDeleted})
}
