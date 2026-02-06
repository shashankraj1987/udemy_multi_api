// Package middlewares provides middleware functions for the API.
package middlewares

import (
	"strings"
	"udemy-multi-api-golang/pkg/response"
	"udemy-multi-api-golang/utils"

	"github.com/gin-gonic/gin"
)

const AuthorizationHeader = "Authorization"

// Authenticate is a middleware that verifies JWT tokens in request headers.
// It extracts the user ID from the token and makes it available in the context.
func Authenticate(c *gin.Context) {
	token := c.GetHeader(AuthorizationHeader)

	if token == "" {
		response.Unauthorized(c, "authorization header is required")
		c.Abort()
		return
	}

	// Remove "Bearer " prefix if present
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	userID, err := utils.VerifyToken(token)
	if err != nil {
		response.Unauthorized(c, "invalid or expired token")
		c.Abort()
		return
	}

	c.Set("userID", userID)
	c.Next()
}
