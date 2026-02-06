package routes

import (
	"udemy-multi-api-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// Events related routes
	authenticated.GET("/events", Get_events)
	authenticated.POST("/events", CreateEvents)
	authenticated.GET("/event_by_id/:id", GetEventById)
	authenticated.PUT("/events/:id", UpdateEvents)
	authenticated.DELETE("/events/:id", DeleteEventById)
	authenticated.POST("/events/:id/register", RegisterForEvents)
	authenticated.DELETE("/events/:id/register", CancelRegisteration)

	// User related routes
	server.POST("/signup", sign_up)
	server.POST("/login", user_Login)

}
