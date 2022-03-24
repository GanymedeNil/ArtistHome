package model

import (
	"gorm.io/gorm"
)

type BlogCategory struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(30);not null;default:\"\";comment:category name;index"`
	Description string `gorm:"column:description;type:varchar(100);not null;default:\"\";comment:category description"`
}
