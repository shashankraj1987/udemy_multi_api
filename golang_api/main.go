// Package main is the entry point of the event registration API.
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"udemy-multi-api-golang/config"
	"udemy-multi-api-golang/db"
	docs "udemy-multi-api-golang/docs"
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

// @title Event Registration API
// @version 1.0
// @description Simple API for managing events and registrations.
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// main initializes the application and starts the HTTP server.
func main() {
	// Load configuration
	cfg := config.Load()

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

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

	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API routes
	routes.RegisterRoutes(engine, db.Client, appLogger)

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
