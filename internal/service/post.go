package service

import (
	"ArtistHome/internal/global"
	"ArtistHome/internal/model"
	"ArtistHome/internal/request"
	"ArtistHome/internal/response"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BlogPost struct {
}

func (p BlogPost) Create(data request.BlogPostForm) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		post := model.BlogPost{Title: data.Title, Summary: data.Summary, Content: data.Content, Status: data.Status}
		result := tx.Create(&post)
		if len(data.Tags) != 0 {
			rows, _ := tx.Model(&model.BlogTag{}).Where("name in (?)", data.Tags).Rows()
			var postTags []model.BlogPostTag
			for rows.Next() {
				var tag model.BlogTag
				tx.ScanRows(rows, &tag)
				postTags = append(postTags, model.BlogPostTag{BlogPostID: post.ID, BlogTagID: tag.ID})
			}
			tx.Create(postTags)
		}
		if len(data.Categories) != 0 {
			rows, _ := tx.Model(&model.BlogCategory{}).Where("name in (?)", data.Categories).Rows()
			var postCategory []model.BlogPostCategory
			for rows.Next() {
				var category model.BlogCategory
				tx.ScanRows(rows, &category)
				postCategory = append(postCategory, model.BlogPostCategory{BlogPostID: post.ID, BlogCategoryID: category.ID})
			}
			tx.Create(postCategory)
		}
		return result.Error
	})
}

func (p BlogPost) Modify(id uint64, data request.BlogPostForm) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		post := model.BlogPost{Title: data.Title, Summary: data.Summary, Content: data.Content, Status: data.Status}
		result := tx.Model(model.BlogPost{}).Where("id = ?", id).Updates(post)
		tx.Where("post_id = ?", post.ID).Delete(&model.BlogPostTag{})
		tx.Where("post_id = ?", post.ID).Delete(&model.BlogPostCategory{})
		if len(data.Tags) != 0 {
			rows, _ := tx.Model(&model.BlogTag{}).Where("name in (?)", data.Tags).Rows()
			var postTags []model.BlogPostTag
			for rows.Next() {
				var tag model.BlogTag
				tx.ScanRows(rows, &tag)
				postTags = append(postTags, model.BlogPostTag{BlogPostID: post.ID, BlogTagID: tag.ID})
			}
			tx.Create(postTags)
		}
		if len(data.Categories) != 0 {
			rows, _ := tx.Model(&model.BlogCategory{}).Where("name in (?)", data.Categories).Rows()
			var postCategory []model.BlogPostCategory
			for rows.Next() {
				var category model.BlogCategory
				tx.ScanRows(rows, &category)
				postCategory = append(postCategory, model.BlogPostCategory{BlogPostID: post.ID, BlogCategoryID: category.ID})
			}
			tx.Create(postCategory)
		}
		return result.Error
	})

}

func (p BlogPost) Delete(id uint64) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&model.BlogPost{}, id)
		tx.Where("blog_post_id = ?", id).Delete(&model.BlogPostTag{})
		tx.Where("blog_post_id = ?", id).Delete(&model.BlogPostCategory{})
		return result.Error
	})

}

func (p BlogPost) Single(id uint64) *response.BlogPostResult {
	var post model.BlogPost
	result := global.DB.Model(&model.BlogPost{}).Where("id = ?", id).Preload(clause.Associations).First(&post)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.LOGGER.Info("post not found#", zap.Uint64("id", id))

		return nil
	}
	res := p.postResponseData(post)
	return &res
}

func (p BlogPost) List(query request.BlogPostQuery) (result []response.BlogPostResult, total int64) {

	postModel := global.DB.Model(&model.BlogPost{})
	if query.Search != "" {
		postModel.Where("title like ?", "%"+query.Search+"%")
	}
	if query.Tag != "" {
		subQuery := global.DB.Model(&model.BlogPostTag{}).Select("blog_post_tags.blog_post_id").
			Joins("left join blog_tags on blog_post_tags.blog_tag_id=blog_tags.id").
			Where("blog_tags.name = ?", query.Tag)
		postModel.Where("id in (?)", subQuery)
	}
	if query.Category != "" {
		subQuery := global.DB.Model(&model.BlogPostCategory{}).Select("blog_post_categories.blog_post_id").
			Joins("left join blog_categories on blog_post_categories.blog_category_id=blog_categories.id").
			Where("blog_categories.name = ?", query.Tag)
		postModel.Where("id in (?)", subQuery)
	}
	err := postModel.Count(&total).Error
	if err != nil {
		return
	}
	var posts []model.BlogPost
	postModel = postModel.Scopes(model.Paginate(query.Page, query.Limit)).
		Preload(clause.Associations).Find(&posts)
	if postModel.Error != nil {
		global.LOGGER.Error("posts find error#" + postModel.Error.Error())
		return
	}
	for _, post := range posts {
		postResult := p.postResponseData(post)
		result = append(result, postResult)
	}
	return
}

func (p BlogPost) postResponseData(post model.BlogPost) response.BlogPostResult {
	return response.BlogPostResult{
		ID:      post.ID,
		Title:   post.Title,
		Summary: post.Summary,
		Content: post.Content,
		Status:  post.Status,
		Tags: func(tags []model.BlogTag) (result []string) {
			for _, tag := range tags {
				result = append(result, tag.Name)
			}
			return
		}(post.Tags),
		Categories: func(categories []model.BlogCategory) (result []string) {
			for _, category := range categories {
				result = append(result, category.Name)
			}
			return
		}(post.Categories),
	}
}
