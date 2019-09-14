package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

// Migrate migrate the database
func Migrate(c *gin.Context) {
	models.AutoMigrate()
}
