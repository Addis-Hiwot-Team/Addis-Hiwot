package http

import (
	"github.com/gin-gonic/gin"

	"addis-hiwot/internal/config"
	"addis-hiwot/internal/delivery/http/handlers"
	"addis-hiwot/internal/repository"
	"addis-hiwot/internal/usecases"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	gormDB := config.NewDB(cfg)
	gormDB.Migrate() // migrated before setting up routes
	db := gormDB.Db

	userRepo := repository.NewUserRepository(db)
	userUC := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUC)

	api := r.Group("/api")
	{
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users", userHandler.GetUsers)
	}
}
