package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sf966gpm/GinLab/routes"
)

func main() {
	r := routes.SetupRoute()

	// Тестовый маршрут
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
