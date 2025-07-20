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
	emailServ := service.NewEmailService()

	otpRepo := repository.NewOtpRepo(db)

	jwtService := service.NewJWTService(cfg.JWTSecret, cfg.TokenDuration)
	userRepo := repository.NewUserRepository(db)
	userUC := usecases.NewUserUsecase(userRepo, jwtService, emailServ, otpRepo)
	userHandler := handlers.NewUserHandler(userUC)

	// daily check-in related
	dailyCheckInRepo := repository.NewDailyCheckInRepository(db)
	dailyCheckInUC := usecases.NewDailyCheckInUsecase(dailyCheckInRepo)
	dailyCheckInHandler := handlers.NewDailyCheckInHandler(dailyCheckInUC)

	sessionRepo := repository.NewSessionRepository(db)
	middleware := middlewares.New(sessionRepo)
	api := r.Group("/api/v1")

	{
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/:id", userHandler.GetUserByID)

		//password related
		api.POST("/users/change_password", middleware.AuthMiddleware(), userHandler.ChangePassword)
		api.POST("/users/forgot_password", userHandler.ForgotPassword)
		api.POST("/users/reset_password", userHandler.ResetPassword)
	}

	auth := api.Group("/auth")
	authRepo := repository.NewAuthRepository(db)
	authUC := usecases.NewAuthUsecase(authRepo, sessionRepo, otpRepo, userRepo, emailServ)
	authHandler := handlers.NewAuthHander(authUC)
	{
		auth.GET("/me", middleware.AuthMiddleware(), authHandler.GetMe)
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/oauth", authHandler.OAuthCodeLoginHandler)
		auth.GET("/activate/:code", authHandler.ActivateAccount)
	}

	// Daily Check-in routes
	checkin := api.Group("/checkin")

	// Apply the existing middleware to all routes in the "checkin" group
	checkin.Use(middleware.AuthMiddleware())
	{
		checkin.POST("", dailyCheckInHandler.AddCheckIn)
		checkin.GET("", dailyCheckInHandler.GetCheckIns)
	}

	api.GET("/protected", middleware.AuthMiddleware(), middlewares.CheckRoles("user"), func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})
}
