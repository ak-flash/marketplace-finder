package controllers

import (
	"github.com/gin-gonic/gin"
	"marketplace-finder/models"
	"net/http"
)

// GET api/getTargets

func GetTargets(c *gin.Context) {
	var targets []models.Target

	targets = models.GetAllTargets()

	c.JSON(http.StatusOK, gin.H{"data": targets})
}
