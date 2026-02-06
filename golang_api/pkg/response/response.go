// Package response provides helpers for returning consistent API responses.
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response structure.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success sends a successful response with optional data.
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error sends an error response with the given status code.
func Error(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// BadRequest sends a 400 Bad Request response.
func BadRequest(c *gin.Context, message string, err string) {
	Error(c, http.StatusBadRequest, message, err)
}

// Unauthorized sends a 401 Unauthorized response.
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, "unauthorized")
}

// Forbidden sends a 403 Forbidden response.
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, message, "forbidden")
}

// NotFound sends a 404 Not Found response.
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, "not_found")
}

// InternalServerError sends a 500 Internal Server Error response.
func InternalServerError(c *gin.Context, message string, err string) {
	Error(c, http.StatusInternalServerError, message, err)
}

// Created sends a 201 Created response with data.
func Created(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusCreated, message, data)
}

// OK sends a 200 OK response with data.
func OK(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusOK, message, data)
}

// NoContent sends a 204 No Content response.
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
