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
// @Summary Register for event
// @Description Register the authenticated user for the target event.
// @Tags Registrations
// @Produce json
// @Param id path int true "Event ID"
// @Security BearerAuth
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events/{id}/registrations [post]
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
// @Summary Unregister from event
// @Description Remove the authenticated user from the event registration list.
// @Tags Registrations
// @Produce json
// @Param id path int true "Event ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events/{id}/registrations [delete]
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
