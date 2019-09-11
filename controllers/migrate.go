package controllers

import (
	"github.com/gin-gonic/gin"
	"katehqbooks/models"
)

func Migrate(c *gin.Context) {
	models.AutoMigrate()
}
