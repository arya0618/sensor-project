package main

import (
	"fmt"
"log"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/arya0618/sensor-project/controllers"
	"github.com/arya0618/sensor-project/models"
	"os"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Connect to PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"))
fmt.Println("=========================================")
fmt.Println(dsn)
	DB, err = gorm.Open("postgres", dsn)
fmt.Println("=========================================")
fmt.Println(err)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}

	// Connect to Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	// Migrate the database
	DB.AutoMigrate(&models.Sensor{})
}

func main() {
	// Initialize Gin router
	router := gin.Default()


// Provide db variable to controllers
    router.Use(func(c *gin.Context) {
        c.Set("db", DB)
        c.Next()
    })
	// Define API routes
	api := router.Group("/api/v1")
	{
		api.GET("/sensors", controllers.GetAllSensors)
		api.POST("/sensor", controllers.CreateSensor)
	}

	// Run the server
	router.Run(":8080")

 log.Println("\n Server started.")
}