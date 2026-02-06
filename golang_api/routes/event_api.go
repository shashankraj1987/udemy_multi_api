// Package routes contains HTTP handlers for event-related operations.
package routes

import (
	"strconv"

	"udemy-multi-api-golang/internal/repository"
	"udemy-multi-api-golang/models"
	"udemy-multi-api-golang/pkg/logger"
	"udemy-multi-api-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

// HandleGetEvents retrieves all events.
// GET /api/events
func HandleGetEvents(eventRepo repository.EventRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		events, err := eventRepo.GetAll()
		if err != nil {
			log.Error("failed to retrieve events", err)
			response.InternalServerError(c, "failed to retrieve events", err.Error())
			return
		}

		response.OK(c, "events retrieved successfully", events)
	}
}

// HandleCreateEvent creates a new event.
// POST /api/events
func HandleCreateEvent(eventRepo repository.EventRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateEventRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to parse create event request", err)
			response.BadRequest(c, "invalid request payload", err.Error())
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			response.Unauthorized(c, "user context not found")
			return
		}

		eventID, err := eventRepo.Create(
			req.Name,
			req.Description,
			req.Location,
			req.DateTime.String(),
			userID.(int64),
		)
		if err != nil {
			log.Error("failed to create event", err)
			response.InternalServerError(c, "failed to create event", err.Error())
			return
		}

		response.Created(c, "event created successfully", gin.H{
			"id":   eventID,
			"name": req.Name,
		})
	}
}

// HandleGetEventByID retrieves a specific event by ID.
// GET /api/events/:id
func HandleGetEventByID(eventRepo repository.EventRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "invalid event ID", err.Error())
			return
		}

		event, err := eventRepo.GetByID(eventID)
		if err != nil {
			log.Error("failed to retrieve event", err)
			response.NotFound(c, "event not found")
			return
		}

		response.OK(c, "event retrieved successfully", event)
	}
}

// HandleUpdateEvent updates an existing event.
// PUT /api/events/:id
func HandleUpdateEvent(eventRepo repository.EventRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "invalid event ID", err.Error())
			return
		}

		var req models.UpdateEventRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to parse update event request", err)
			response.BadRequest(c, "invalid request payload", err.Error())
			return
		}

		// Retrieve the event to verify ownership
		event, err := eventRepo.GetByID(eventID)
		if err != nil {
			log.Error("failed to retrieve event", err)
			response.NotFound(c, "event not found")
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			response.Unauthorized(c, "user context not found")
			return
		}

		// Check if the user owns the event
		if event["userId"].(int64) != userID.(int64) {
			response.Forbidden(c, "you do not have permission to update this event")
			return
		}

		// Update the event
		err = eventRepo.Update(eventID, req.Name, req.Description, req.Location, req.DateTime.String())
		if err != nil {
			log.Error("failed to update event", err)
			response.InternalServerError(c, "failed to update event", err.Error())
			return
		}

		response.OK(c, "event updated successfully", gin.H{"id": eventID})
	}
}

// HandleDeleteEvent deletes an event.
// DELETE /api/events/:id
func HandleDeleteEvent(eventRepo repository.EventRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "invalid event ID", err.Error())
			return
		}

		// Retrieve the event to verify ownership
		event, err := eventRepo.GetByID(eventID)
		if err != nil {
			log.Error("failed to retrieve event for deletion", err)
			response.NotFound(c, "event not found")
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			response.Unauthorized(c, "user context not found")
			return
		}

		// Check if the user owns the event
		if event["userId"].(int64) != userID.(int64) {
			response.Forbidden(c, "you do not have permission to delete this event")
			return
		}

		// Delete the event
		rowsAffected, err := eventRepo.Delete(eventID)
		if err != nil {
			log.Error("failed to delete event", err)
			response.InternalServerError(c, "failed to delete event", err.Error())
			return
		}

		response.OK(c, "event deleted successfully", gin.H{
			"id":            eventID,
			"rows_affected": rowsAffected,
		})
	}
}
