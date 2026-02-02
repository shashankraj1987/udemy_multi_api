package middlewares

import (
	"net/http"
	"udemy-multi-api-golang/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	c.Set("UId", userId)

	c.Next()
}
