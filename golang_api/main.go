// Package main is the entry point of the event registration API.
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"udemy-multi-api-golang/config"
	"udemy-multi-api-golang/db"
	"udemy-multi-api-golang/pkg/logger"
	"udemy-multi-api-golang/routes"
	"udemy-multi-api-golang/utils"
)

func init() {
	// Set Gin mode based on environment
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}
}

// main initializes the application and starts the HTTP server.
func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	appLogger := logger.New(cfg.Logging.Level)
	appLogger.Info("starting event registration api")

	// Initialize database
	err := db.InitDB(cfg)
	if err != nil {
		appLogger.Error("failed to initialize database", err)
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Initialize JWT configuration
	utils.InitJWT(&cfg.JWT)

	// Create Gin engine
	engine := gin.Default()

	// Set up trusted proxies
	engine.SetTrustedProxies([]string{"0.0.0.0"})

	// Register API routes
	routes.RegisterRoutes(engine, db.DB, appLogger)

	// Health check endpoint
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port
	appLogger.Info("starting http server on " + serverAddr)

	if err := engine.Run(serverAddr); err != nil {
		appLogger.Error("failed to start server", err)
		log.Fatal(err)
	}
}
