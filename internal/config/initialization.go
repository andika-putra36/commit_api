package config

import (
	"commit_api/internal/task"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeEverything() *gin.Engine {
	db := initializeDB()
	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := task.NewHandler(taskService)

	router := gin.Default()
	v1 := router.Group("api/v1")

	task.RegisterRoutes(v1, *taskHandler)
	return router
	// router.Run(":8888")
}

func initializeDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error getting .env file")
	}

	env := getEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.Host,
		env.User,
		env.Password,
		env.Name,
		env.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database is failed to connect.")
	}

	return db
}

func getEnv() Env {
	return Env{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}
