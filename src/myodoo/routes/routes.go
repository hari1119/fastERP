package routes

import (
	"github.com/gin-gonic/gin"
	"myodoo/handlers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/users", handlers.CreateUserHandler)
	router.GET("/users/:id", handlers.GetUserHandler)
        router.GET("/", func(c *gin.Context) {c.JSON(200, gin.H{"message": "Welcome to Go-Odoo"})})

}
