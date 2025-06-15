package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"addis-hiwot/internal/config"
	"addis-hiwot/internal/delivery/http/handlers"
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/repository"
	"addis-hiwot/internal/usecases"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto-migrate your models
	db.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(db)
	userUC := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUC)

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.GetUsers)
	}
}
