package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase(cfg *AppConfig) (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db.Debug(), nil
}

func ConnectDatabaseWithRetryIn60s(cfg *AppConfig) (*gorm.DB, error) {
	const timeRetry = time.Second * 60
	var db *gorm.DB
	var err error
	deadline := time.Now().Add(timeRetry)

	for time.Now().Before(deadline) {
		db, err = connectDatabase(cfg)
		if err == nil {
			return db, nil
		}
		time.Sleep(time.Second * 5)
	}

	return nil, fmt.Errorf("failed to connect to database: %w", err)
}
