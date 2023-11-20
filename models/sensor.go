// models/sensor.go
package models

import (
  "fmt"
"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB is a global variable to be used across packages for database connection.
var DB *gorm.DB

// Sensor model represents a sensor with parameters.
type Sensor struct {
	gorm.Model
	Codename    string  `json:"codename"`
	Coordinates string  `json:"coordinates"`
	DataRate    float64 `json:"datarate"`
}

// Initialize initializes the database connection.
func Initialize(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", databaseURL)
	if err != nil {
fmt.Println("error-",err)
		return nil, err
	}

	DB = db

	// AutoMigrate will attempt to automatically migrate the schema,
	// if it's not already in sync with the registered models.
	//DB.AutoMigrate(&Sensor{})


if err := DB.AutoMigrate(&Sensor{}).Error; err != nil {
        log.Printf("Failed to auto-migrate the database: %v", err)
        return nil, err
    }

	return DB, nil
}
