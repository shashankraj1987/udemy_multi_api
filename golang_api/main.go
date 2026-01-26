package main

import (
	"fmt"
	"udemy-multi-api-golang/db"
	"udemy-multi-api-golang/models"
	"udemy-multi-api-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.SetTrustedProxies([]string{"0.0.0.0"})
	db.InitDb()

	fmt.Println("Available models:")
	for _, model := range models.GetModels() {
		fmt.Printf("- %T\n", model)
	}

	routes.RegisterRoutes(server)
	server.Run(":8081")
}
