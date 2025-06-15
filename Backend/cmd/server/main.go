package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"addis-hiwot/internal/config"
	"addis-hiwot/internal/delivery/http"
)

func main() {
	err := godotenv.Load()
	log.Println("go env", err)
	if err != nil {
		log.Println("No .env file found")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config error:", err)
	}

	router := gin.Default()
	http.SetupRoutes(router, cfg)

	log.Println("Server running at", cfg.ServerPort)
	router.Run(cfg.ServerPort)
}
