package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

func Make() (*gorm.DB, error) {
	return NewWithDefault(DSN())
}

func NewWithDefault(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	SetDefault(db)
	return db, nil
}

func NewFromDSN(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn))
}

func Raw(sql string, values ...any) *gorm.DB {
	return Q.db.Raw(sql, values...)
}
