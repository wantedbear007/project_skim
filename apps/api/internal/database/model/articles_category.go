package model

import (
	"gorm.io/gorm"
)


type ArticleCategory struct {

	gorm.Model

	Name 									string							`gorm:"size:100;not null;uniqueIndex;"`
	LogoURL								string							`gorm:"size:255"`
	Description						string							`gorm:"type:text"`

	Articles						[]Articles						`gorm:"foreignKey:CategoryID"`
}