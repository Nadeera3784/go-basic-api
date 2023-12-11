
package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"application/src/models"
)

type MySQLConfig struct {
	Port     string
	Host     string
	Database string
	Username string
	Password string
}

var MySQLConfigInstance MySQLConfig

var Instance *gorm.DB;

// SetupDatabase initializes the database connection and performs auto-migration
func SetupDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	MySQLConfigInstance.Port = os.Getenv("DB_PORT")
	MySQLConfigInstance.Host = os.Getenv("DB_HOST")
	MySQLConfigInstance.Database = os.Getenv("DB_DATABASE")
	MySQLConfigInstance.Username = os.Getenv("DB_USERNAME")
	MySQLConfigInstance.Password = os.Getenv("DB_PASSWORD")

	db, err := Connect()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	// Auto-migrate models
	err = AutoMigrate(db);
	if err != nil {
		return nil, fmt.Errorf("error performing auto-migration: %v", err)
	}

	Instance = db;

	return db, nil
}

// Connect establishes a database connection
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MySQLConfigInstance.Username,
		MySQLConfigInstance.Password,
		MySQLConfigInstance.Host,
		MySQLConfigInstance.Port,
		MySQLConfigInstance.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening the database connection: %v", err)
	}

	return db, nil
}

// AutoMigrate performs auto-migration for models
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		return fmt.Errorf("error auto-migrating models: %v", err)
	}
	return nil
}
