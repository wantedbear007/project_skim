package migrations

import (
	// "context"
	"fmt"

	"github.com/wantedbear007/project_skim/internal/database/model"
	"github.com/wantedbear007/project_skim/internal/env"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func RunMigrations() {
	db_url := env.DATABASE_URL.GetValue()

	zap.L().Info("Migration started")

	db, error := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if error != nil {
		fmt.Println("Failed to connect with database ")
		panic(error)
	}

	// to migrate schema 
	db.AutoMigrate(&model.ArticleCategory{})
	db.AutoMigrate(&model.Articles{})
	db.AutoMigrate(&model.ArticlesSource{})

	zap.L().Info("sucessfully ran migrations")



}

