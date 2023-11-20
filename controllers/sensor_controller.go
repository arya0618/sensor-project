package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/arya0618/sensor-project/models"
//	"gorm.io/gorm"
"fmt"
)

// GetAllSensors retrieves all sensors from the database.
func GetAllSensors(c *gin.Context) {
	var sensors []models.Sensor
	models.DB.Find(&sensors)
	c.JSON(200, sensors)
}

// CreateSensor creates a new sensor in the database.
func CreateSensor(c *gin.Context) {
fmt.Println("---------in controller--------")
	var input models.Sensor
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
fmt.Println("--------- 2 log in controller--------",models.DB)
	// Validate and save to the database
	if err := models.DB.Create(&input).Error; err != nil {
		fmt.Println("Error creating sensor:", err)
		c.JSON(500, gin.H{"error": "Failed to create sensor"})
		return
	}
fmt.Println("--------- 3 log in controller--------")
	c.JSON(200, input)
}

