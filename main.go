package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zombox0633/printer_backend_go/src"
	"github.com/zombox0633/printer_backend_go/src/config"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Warning: .env file not found 🙀")
	}

	env := config.LoadConfig()
	if env.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	src.RoutersGroup(router)

	log.Printf("Server starting on http://localhost:%s 🎉", env.Port)

	if err := router.Run(":" + env.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
