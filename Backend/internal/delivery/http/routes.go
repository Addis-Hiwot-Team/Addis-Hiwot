package http

import (
	"github.com/gin-gonic/gin"

	"addis-hiwot/internal/config"
	"addis-hiwot/internal/delivery/http/handlers"
	"addis-hiwot/internal/delivery/middlewares"
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

	sessionRepo := repository.NewSessionRepository(db)
	middleware := middlewares.New(sessionRepo)
	api := r.Group("/api/v1")

	{
		api.GET("/users", userHandler.GetUsers)
	}

	auth := api.Group("/auth")
	authRepo := repository.NewAuthRepository(db)
	authUC := usecases.NewAuthUsecase(authRepo, sessionRepo)
	authHandler := handlers.NewAuthHander(authUC)
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
		auth.POST("/refresh", authHandler.Refresh)
	}
	api.GET("/protected", middleware.AuthMiddleware(), middlewares.CheckRoles("user"), func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})
}
