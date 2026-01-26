package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Events related routes
	server.GET("/events", Get_events)
	server.POST("/events", CreateEvents)
	server.GET("/event_by_id/:id", GetEventById)
	server.PUT("/events/:id", UpdateEvents)
	server.DELETE("/events/:id", DeleteEventById)

	// User related routes
	server.POST("/signup", signup)
}
