package model

import (
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title      string         `gorm:"column:title;type:varchar(30);not null;default:\"\";comment:post title"`
	Summary    string         `gorm:"column:summary;type:text;size:255;not null;comment:post summary"`
	Content    string         `gorm:"column:content;type:json;comment:post content"`
	Status     uint           `gorm:"column:status;type:tinyint(1);not null;default:0;comment:post status  1:draft 2:pending 3:private 4:publish"`
	AuthorID   uint           `gorm:"column:author_id;type:bigint(20);not null;default:0;comment:post author;index"`
	Tags       []BlogTag      `gorm:"many2many:blog_post_tags;"`
	Categories []BlogCategory `gorm:"many2many:blog_post_categories;"`
}
