package config

import "os"

var (
	Port       = os.Getenv("PORT")
	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
	JwtSecret  = os.Getenv("JWT_SECRET")
)
