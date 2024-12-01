package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	migrateDown := os.Getenv("MIGRATE_DOWN") == "true"

	if err := runMigrations(sqlDB, migrateDown); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

func runMigrations(db *sql.DB, migrateDown bool) error {
	if migrateDown {
		if err := execSQLFile(db, "migrations/down.sql"); err != nil {
			return fmt.Errorf("failed to run down migration: %w", err)
		}
	}
	if err := execSQLFile(db, "migrations/up.sql"); err != nil {
		return fmt.Errorf("failed to run up migration: %w", err)
	}
	return nil
}

func execSQLFile(db *sql.DB, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	return err
}
