package model

type BlogPostTag struct {
	ID         uint `gorm:"primarykey"`
	BlogPostID uint `gorm:"column:blog_post_id;type:bigint(20);not null;default:0;comment:post id;index"`
	BlogTagID  uint `gorm:"column:blog_tag_id;type:bigint(20);not null;default:0;comment:tag id;index"`
}
