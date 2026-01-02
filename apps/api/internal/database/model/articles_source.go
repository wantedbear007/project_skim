package model

import (
	"gorm.io/gorm"
)

type ArticlesSource struct {

	gorm.Model

	SourceTitle						string						`gorm:"not null;uniqueIndex"`
	SourceHomePage				string						`gorm:"uniqueIndex"`

	Articles						[]Articles					`gorm:"foreignKey:SourceID"`
}

