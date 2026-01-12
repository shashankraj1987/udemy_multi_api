package main

import (
	"net/http"
	"udemy-golang-api/db"
	"udemy-golang-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDb()

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse the request to Json."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created", "event": event})
}
