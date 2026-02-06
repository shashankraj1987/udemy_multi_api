// Package db provides database initialization and connection management.
package db

import (
	"log"
	"time"

	"udemy-multi-api-golang/config"
	"udemy-multi-api-golang/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Client is the global GORM database handle.
var Client *gorm.DB

// InitDB initializes the database connection using GORM and runs migrations.
func InitDB(cfg *config.Config) error {
	database, err := gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.MaxConnLifetime) * time.Second)

	if err := database.AutoMigrate(&models.User{}, &models.Event{}, &models.Registration{}); err != nil {
		log.Printf("failed to run database migrations: %v\n", err)
		return err
	}

	Client = database
	log.Printf("database connected successfully: %s\n", cfg.Database.Path)
	return nil
}

// CloseDB closes the underlying database connection.
func CloseDB() error {
	if Client == nil {
		return nil
	}

	sqlDB, err := Client.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
