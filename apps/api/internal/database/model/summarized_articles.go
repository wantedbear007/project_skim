package model

import (
	"gorm.io/gorm"
)

type Articles struct {

	gorm.Model

	Title									string						`gorm:"not null"`
	ArticleUrl						string						`gorm:"uniqueIndex"`
	Description						string
	ImageUrl							string
	PubDate								string


	SourceID							uint
	ArticleSource					ArticlesSource		`gorm:"foreignKey:SourceID;references:ID"`

	CategoryID						uint
	ArticleCategory				ArticleCategory		`gorm:"foreignKey:CategoryID;references:ID"`
}

