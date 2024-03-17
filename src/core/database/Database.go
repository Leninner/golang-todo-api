package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"todo-api/src/tasks"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
}

func LoadConfig() DatabaseConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}

	return DatabaseConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Dbname:   os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func BuildDatabaseString(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.Host, config.Port, config.User, config.Dbname, config.Password)
}

func GetConnection() {
	config := LoadConfig()
	databaseString := BuildDatabaseString(config)

	var err error
	DB, err = gorm.Open(postgres.Open(databaseString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connected")
}

func SetupModels() error {
	err := DB.AutoMigrate(&tasks.Task{})

	if err != nil {
		return errors.New("error migrating models")
	}

	log.Println("Models migrated successfully")
	return nil
}
