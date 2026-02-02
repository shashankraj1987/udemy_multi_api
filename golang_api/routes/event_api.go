package routes

import (
	"net/http"
	"strconv"
	"udemy-multi-api-golang/models"
	"udemy-multi-api-golang/utils"

	"github.com/gin-gonic/gin"
)

func Get_events(context *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching Data from the Database", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)

}

func CreateEvents(context *gin.Context) {

	var err error

	// Validate tokens via Auth Headers.
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized",
			"error": "Empty Auth Token."})
		return
	}

	// Check the token to see if it is valid.

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized",
			"error": err.Error()})
		return
	}

	// If Valid, accept the token and move forward.
	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse the value.", "event": event})
		return
	}
	// event.ID = 1
	event.USerID = int(userId)
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Saving Data in the Database", "Error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "event Created",
		"event": event})
}

func GetEventById(context *gin.Context) {
	event_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect ID passed as Integer", "Error": err.Error()})
		return
	}
	event, err := models.GetId(event_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Unable to get event", "error": err.Error()})
	}

	context.JSON(http.StatusOK, event)
}

func UpdateEvents(context *gin.Context) {
	event_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect ID passed as Integer", "Error": err.Error()})
		return
	}

	_, err = models.GetId(event_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Unable to get event", "error": err.Error()})
	}

	var updatedEvent models.Event

	err = context.ShouldBind(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse the value.", "event_id": event_id})
		return
	}

	updatedEvent.ID = event_id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Message": "Event Updated Successfully."})
}

func DeleteEventById(context *gin.Context) {
	event_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID", "Error": err.Error()})
		return
	}

	event, err := models.GetId(event_id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Id Not found in the Database",
			"error":   err.Error()})
		return
	}

	rows, err := event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to Delete ID.",
			"error":   err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":       "Event Deleted",
		"rows effected": rows})
}
