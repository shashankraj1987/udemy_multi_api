// Package routes registers all API routes and their corresponding handlers.
package routes

import (
	"database/sql"

	"udemy-multi-api-golang/internal/repository"
	"udemy-multi-api-golang/middlewares"
	"udemy-multi-api-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all API routes with the Gin engine.
func RegisterRoutes(engine *gin.Engine, db *sql.DB, log logger.Logger) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository(db)
	regRepo := repository.NewRegistrationRepository(db)

	// Public routes (no authentication required)
	auth := engine.Group("/auth")
	{
		auth.POST("/signup", HandleSignUp(userRepo, log))
		auth.POST("/login", HandleLogin(userRepo, log))
	}

	// Protected routes (authentication required)
	api := engine.Group("/api")
	api.Use(middlewares.Authenticate)
	{
		// Event routes
		events := api.Group("/events")
		{
			events.GET("", HandleGetEvents(eventRepo, log))
			events.POST("", HandleCreateEvent(eventRepo, log))
			events.GET("/:id", HandleGetEventByID(eventRepo, log))
			events.PUT("/:id", HandleUpdateEvent(eventRepo, log))
			events.DELETE("/:id", HandleDeleteEvent(eventRepo, log))
		}

		// Event registration routes
		registrations := api.Group("/events/:id/registrations")
		{
			registrations.POST("", HandleRegisterEvent(eventRepo, regRepo, log))
			registrations.DELETE("", HandleUnregisterEvent(regRepo, log))
		}
	}
}
