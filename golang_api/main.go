package main

import (
	"net/http"
	"udemy-multi-api-golang/db"
	"udemy-multi-api-golang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	db.InitDb()

	server.GET("/events", Get_events)
	server.POST("/events", CreateEvents)
	server.Run(":8081")
}

func Get_events(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)

}

func CreateEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Could not parse the value.", "event": event})
		return
	}
	event.ID = 1
	event.USerID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"Message": "event Created",
		"event": event})
}
