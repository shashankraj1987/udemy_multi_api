// Package routes contains HTTP handlers for event registration operations.
package routes

import (
	"strconv"

	"udemy-multi-api-golang/internal/repository"
	"udemy-multi-api-golang/pkg/logger"
	"udemy-multi-api-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

// HandleRegisterEvent registers a user for an event.
// POST /api/events/:id/registrations
func HandleRegisterEvent(
	eventRepo repository.EventRepository,
	regRepo repository.RegistrationRepository,
	log logger.Logger,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "invalid event ID", err.Error())
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			response.Unauthorized(c, "user context not found")
			return
		}

		// Check if event exists
		_, err = eventRepo.GetByID(eventID)
		if err != nil {
			log.Error("failed to retrieve event for registration", err)
			response.NotFound(c, "event not found")
			return
		}

		// Check if already registered
		isRegistered, err := regRepo.IsRegistered(eventID, userID.(int64))
		if err != nil {
			log.Error("failed to check registration status", err)
			response.InternalServerError(c, "failed to process registration", err.Error())
			return
		}

		if isRegistered {
			response.BadRequest(c, "user is already registered for this event", "duplicate registration")
			return
		}

		// Register user for event
		err = regRepo.Register(eventID, userID.(int64))
		if err != nil {
			log.Error("failed to register user for event", err)
			response.InternalServerError(c, "failed to register for event", err.Error())
			return
		}

		response.Created(c, "successfully registered for event", gin.H{
			"event_id": eventID,
			"user_id":  userID,
		})
	}
}

// HandleUnregisterEvent unregisters a user from an event.
// DELETE /api/events/:id/registrations
func HandleUnregisterEvent(regRepo repository.RegistrationRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "invalid event ID", err.Error())
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			response.Unauthorized(c, "user context not found")
			return
		}

		// Check if user is registered
		isRegistered, err := regRepo.IsRegistered(eventID, userID.(int64))
		if err != nil {
			log.Error("failed to check registration status", err)
			response.InternalServerError(c, "failed to process unregistration", err.Error())
			return
		}

		if !isRegistered {
			response.BadRequest(c, "user is not registered for this event", "not registered")
			return
		}

		// Unregister user from event
		err = regRepo.Unregister(eventID, userID.(int64))
		if err != nil {
			log.Error("failed to unregister user from event", err)
			response.InternalServerError(c, "failed to unregister from event", err.Error())
			return
		}

		response.OK(c, "successfully unregistered from event", gin.H{
			"event_id": eventID,
			"user_id":  userID,
		})
	}
}
