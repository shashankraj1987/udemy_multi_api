package response
// Package response provides standardized response structures for the API.
// It ensures consistent response formatting across all endpoints.
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
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





















































}	c.JSON(http.StatusNoContent, nil)func NoContent(c *gin.Context) {// NoContent sends a 204 No Content response.}	Success(c, http.StatusOK, message, data)func OK(c *gin.Context, message string, data interface{}) {// OK sends a 200 OK response with data.}	Success(c, http.StatusCreated, message, data)func Created(c *gin.Context, message string, data interface{}) {// Created sends a 201 Created response with data.}	Error(c, http.StatusInternalServerError, message, err)func InternalServerError(c *gin.Context, message string, err string) {// InternalServerError sends a 500 Internal Server Error response.}	Error(c, http.StatusNotFound, message, "not_found")func NotFound(c *gin.Context, message string) {// NotFound sends a 404 Not Found response.}	Error(c, http.StatusForbidden, message, "forbidden")func Forbidden(c *gin.Context, message string) {// Forbidden sends a 403 Forbidden response.}	Error(c, http.StatusUnauthorized, message, "unauthorized")func Unauthorized(c *gin.Context, message string) {// Unauthorized sends a 401 Unauthorized response.}	Error(c, http.StatusBadRequest, message, err)func BadRequest(c *gin.Context, message string, err string) {// BadRequest sends a 400 Bad Request response.}	})		Error:   err,		Message: message,		Success: false,	c.JSON(statusCode, Response{func Error(c *gin.Context, statusCode int, message string, err string) {// Error sends an error response.}	})		Data:    data,		Message: message,