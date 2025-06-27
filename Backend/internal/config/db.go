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
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	var db *gorm.DB
	var err error

	maxAttempts := 5
	for i := 1; i <= maxAttempts; i++ {
		db, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
			Logger: gormLogger,
		})

		if err == nil {
			log.Printf("✅ Connected to DB on attempt %d", i)
			break
		}

		log.Printf("❌ Attempt %d: failed to connect to DB: %v", i, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("❌ All attempts to connect to DB failed: %v", err)
	}

	return &DB{Db: db}
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
