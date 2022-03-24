package model

import (
	"gorm.io/gorm"
)

type BlogTag struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(30);not null;default:\"\";comment:tag name;index"`
	Description string `gorm:"column:description;type:varchar(100);not null;default:\"\";comment:tag description"`
}
