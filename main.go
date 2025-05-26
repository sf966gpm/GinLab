package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализируем Gin-роутер
	r := gin.Default()

	// Роут для проверки работы сервера
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to GinLab!",
			"version": "1.0.2",
		})
	})

	// Роут для health-check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Роут с параметром
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	// Запускаем сервер на порту 8080
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
