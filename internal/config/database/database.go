package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	gormpg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() (*gorm.DB, error) {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Connect database
	db, err := gorm.Open(gormpg.Open(getDSN()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	// Run migrations
	err = RunMigrations()
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	DB = db
	return db, nil
}

func RunMigrations() error {
	db, err := sql.Open("postgres", getDSN())
	if err != nil {
		return fmt.Errorf("failed to open db: %w", err)
	}
	defer db.Close()

	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate: %w", err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run up: %w", err)
	}

	log.Println("Migrations completed")
	return nil
}

func getDSN() string {
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return dsn
	}
	return "postgres://postgres:password@localhost:5432/elearning?sslmode=disable"
}

func GetDB() *gorm.DB {
	return DB
}
