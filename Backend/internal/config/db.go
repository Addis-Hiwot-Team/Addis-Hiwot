package config

import (
	"addis-hiwot/internal/domain/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Db *gorm.DB
}

func NewDB(cfg *Config) *DB {

	gromLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore record not found error
			Colorful:                  true,         // Enable colorful output
		},
	)
	db, err := gorm.Open(
		postgres.Open(cfg.DSN()),
		&gorm.Config{
			Logger: gromLogger,
		},
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &DB{
		Db: db,
	}
}

func (db *DB) Migrate() {
	if err := db.Db.AutoMigrate(
		&models.User{},
		&models.AIChatHistory{},
		&models.CommunityMessage{},
		&models.DailyCheckIn{},
		&models.GoalReward{},
		&models.Goal{},
		&models.HabitLog{},
		&models.Habit{},
		&models.MotivationalQuote{},
		&models.Resource{},
		&models.UserQuote{},
		&models.UserResource{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migration completed successfully")
}
