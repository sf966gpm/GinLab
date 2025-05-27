package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	// Логика получения пользователей
	c.JSON(http.StatusOK, gin.H{"data": "список пользователей"})
}

func CreateUser(c *gin.Context) {
	// Логика создания пользователя
	c.JSON(http.StatusCreated, gin.H{"data": "пользователь создан"})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	// Логика получения пользователя по ID
	c.JSON(http.StatusOK, gin.H{"data": "пользователь " + id})
}
