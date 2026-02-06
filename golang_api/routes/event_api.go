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
// @Summary List events
// @Description Retrieve all available events.
// @Tags Events
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events [get]
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
// @Summary Create event
// @Description Create a new event owned by the authenticated user.
// @Tags Events
// @Accept json
// @Produce json
// @Param payload body models.CreateEventRequest true "Event payload"
// @Security BearerAuth
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events [post]
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

		event := &models.Event{
			Name:        req.Name,
			Description: req.Description,
			Location:    req.Location,
			DateTime:    req.DateTime,
			UserID:      userID.(int64),
		}

		if err := eventRepo.Create(event); err != nil {
			log.Error("failed to create event", err)
			response.InternalServerError(c, "failed to create event", err.Error())
			return
		}

		response.Created(c, "event created successfully", gin.H{
			"id":   event.ID,
			"name": event.Name,
		})
	}
}

// HandleGetEventByID retrieves a specific event by ID.
// @Summary Get event
// @Description Retrieve event details by ID.
// @Tags Events
// @Produce json
// @Param id path int true "Event ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events/{id} [get]
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
// @Summary Update event
// @Description Update an event owned by the authenticated user.
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param payload body models.UpdateEventRequest true "Event payload"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events/{id} [put]
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
		if event.UserID != userID.(int64) {
			response.Forbidden(c, "you do not have permission to update this event")
			return
		}

		// Update the event
		event.Name = req.Name
		event.Description = req.Description
		event.Location = req.Location
		event.DateTime = req.DateTime

		if err := eventRepo.Update(event); err != nil {
			log.Error("failed to update event", err)
			response.InternalServerError(c, "failed to update event", err.Error())
			return
		}

		response.OK(c, "event updated successfully", gin.H{"id": eventID})
	}
}

// HandleDeleteEvent deletes an event.
// @Summary Delete event
// @Description Delete an event owned by the authenticated user.
// @Tags Events
// @Produce json
// @Param id path int true "Event ID"
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 403 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/events/{id} [delete]
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
		if event.UserID != userID.(int64) {
			response.Forbidden(c, "you do not have permission to delete this event")
			return
		}

		// Delete the event
		if err := eventRepo.Delete(eventID); err != nil {
			log.Error("failed to delete event", err)
			response.InternalServerError(c, "failed to delete event", err.Error())
			return
		}

		response.OK(c, "event deleted successfully", gin.H{
			"id": eventID,
		})
	}
}
