package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	ServerPort    string
	JWTSecret     string
	TokenDuration time.Duration
}

func LoadConfig() (*Config, error) {
	durationStr := os.Getenv("TOKEN_DURATION")
	durationHours, err := strconv.Atoi(durationStr)
	if err != nil || durationHours <= 0 {
		durationHours = 24 // default fallback
	}
	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		ServerPort:    os.Getenv("SERVER_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		TokenDuration: time.Duration(durationHours) * time.Hour,
	}, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s channel_binding=require sslmode=require TimeZone=Africa/Addis_Ababa",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
}
