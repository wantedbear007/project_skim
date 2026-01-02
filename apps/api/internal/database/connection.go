package database

import (
	"github.com/wantedbear007/project_skim/internal/env"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db_url := env.DATABASE_URL.GetValue()

	db, error := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if error != nil {
		zap.L().Error("Failed to connect to database", zap.Error(error))

		return nil, error
	}

	return db, nil

}


