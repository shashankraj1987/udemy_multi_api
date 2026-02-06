package routes

import (
	"net/http"
	"strconv"
	"udemy-multi-api-golang/models"

	"github.com/gin-gonic/gin"
)

func RegisterForEvents(context *gin.Context) {
	userID := context.GetInt64("UId")
	eventId, err := strconv.ParseInt(context.Param("Id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetId(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
	}
	event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event.", "error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func CancelRegisteration(context *gin.Context) error {
	userId := context.GetInt64("UId")
	eventId, err := strconv.ParseInt(context.Param("Id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return err
	}

	var event models.Event
	event.ID = eventId

	err = event.UnRegister(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not Cancel the event."})
		return err
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Registeration Cancelled."})

	return nil

}
