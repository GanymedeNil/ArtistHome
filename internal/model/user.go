package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(80);not null;default:\"\";comment:user name;index"`
	Password string `gorm:"column:password;type:varchar(60);not null;default:\"\";comment:user password hash"`
}
