package main

import (
	"log"

	"addis-hiwot/docs"
	"addis-hiwot/internal/config"
	"addis-hiwot/internal/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Addis Hiwot API
// @version         1.0
// @description     This is api documentaion for AddisHiwot API.

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
func main() {
	err := godotenv.Load()
	log.Println("go env", err)
	if err != nil {
		log.Println("No .env file found")
	}
	docs.SwaggerInfo.BasePath = "/api"
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config error:", err)
	}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	http.SetupRoutes(router, cfg)

	port := cfg.ServerPort
	if port == "" {
		port = "8080"
	}
	if port[0] != ':' {
		port = ":" + port
	}

	log.Println("Server running at", port)
	router.Run(port)
}
