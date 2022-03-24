package model

type BlogPostCategory struct {
	ID             uint `gorm:"primarykey"`
	BlogPostID     uint `gorm:"column:blog_post_id;type:bigint(20);not null;;default:0;comment:post id;index"`
	BlogCategoryID uint `gorm:"column:blog_category_id;type:bigint(20);not null;default:0;comment:category id;index"`
}
