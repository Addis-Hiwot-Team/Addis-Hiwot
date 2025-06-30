package http

import (
	"github.com/gin-gonic/gin"

	"addis-hiwot/internal/config"
	"addis-hiwot/internal/delivery/http/handlers"
	"addis-hiwot/internal/repository"
	"addis-hiwot/internal/service"
	"addis-hiwot/internal/usecases"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	gormDB := config.NewDB(cfg)
	gormDB.Migrate() // migrated before setting up routes
	db := gormDB.Db

	jwtService := service.NewJWTService(cfg.JWTSecret, cfg.TokenDuration)
	userRepo := repository.NewUserRepository(db)
	userUC := usecases.NewUserUsecase(userRepo, jwtService)
	userHandler := handlers.NewUserHandler(userUC)

	api := r.Group("/api/v1")

	// AUthentication routes
	auth := api.Group("/auth")
	{
		auth.POST("/register", userHandler.CreateUser)
		auth.POST("/login", userHandler.LoginUser)
	}

	{
		api.GET("/users", userHandler.GetUsers)
	}
}
