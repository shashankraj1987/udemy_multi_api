// Package routes contains HTTP handlers for user-related operations.
package routes

import (
	"udemy-multi-api-golang/internal/repository"
	"udemy-multi-api-golang/models"
	"udemy-multi-api-golang/pkg/logger"
	"udemy-multi-api-golang/pkg/response"
	"udemy-multi-api-golang/utils"

	"github.com/gin-gonic/gin"
)

// HandleSignUp handles user account creation.
// @Summary User signup
// @Description Create a new user account with email and password.
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body models.SignupRequest true "Signup request"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /auth/signup [post]
func HandleSignUp(userRepo repository.UserRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.SignupRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to parse signup request", err)
			response.BadRequest(c, "invalid request payload", err.Error())
			return
		}

		// Check if user already exists
		exists, err := userRepo.Exists(req.Email)
		if err != nil {
			log.Error("failed to check user existence", err)
			response.InternalServerError(c, "failed to process request", err.Error())
			return
		}

		if exists {
			response.BadRequest(c, "user already exists", "email already registered")
			return
		}

		// Hash the password
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			log.Error("failed to hash password", err)
			response.InternalServerError(c, "failed to create user", err.Error())
			return
		}

		// Create the user
		userID, err := userRepo.Create(req.Email, hashedPassword)
		if err != nil {
			log.Error("failed to create user", err)
			response.InternalServerError(c, "failed to create user", err.Error())
			return
		}

		response.Created(c, "user created successfully", gin.H{
			"id":    userID,
			"email": req.Email,
		})
	}
}

// HandleLogin handles user authentication.
// @Summary User login
// @Description Authenticate a user and receive a JWT token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param payload body models.LoginRequest true "Login request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /auth/login [post]
func HandleLogin(userRepo repository.UserRepository, log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			log.Error("failed to parse login request", err)
			response.BadRequest(c, "invalid request payload", err.Error())
			return
		}

		// Get user from repository
		userID, passwordHash, err := userRepo.GetByEmail(req.Email)
		if err != nil {
			log.Warn("login attempt for non-existent user", req.Email)
			response.Unauthorized(c, "invalid email or password")
			return
		}

		// Check password
		if !utils.CheckPassword(req.Password, passwordHash) {
			log.Warn("login attempt with invalid password", req.Email)
			response.Unauthorized(c, "invalid email or password")
			return
		}

		// Generate token
		token, err := utils.GenerateToken(req.Email, userID)
		if err != nil {
			log.Error("failed to generate token", err)
			response.InternalServerError(c, "failed to authenticate user", err.Error())
			return
		}

		response.OK(c, "login successful", gin.H{
			"token": token,
			"user": gin.H{
				"id":    userID,
				"email": req.Email,
			},
		})
	}
}
