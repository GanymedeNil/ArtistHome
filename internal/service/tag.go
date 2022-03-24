package service

import (
	"ArtistHome/internal/global"
	"ArtistHome/internal/model"
	"ArtistHome/internal/request"
	"ArtistHome/internal/response"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BlogTag struct {
}

func (t BlogTag) Create(data request.BlogTagForm) error {
	tag := model.BlogTag{
		Name:        data.Name,
		Description: data.Description,
	}
	return global.DB.Create(&tag).Error

}

func (t BlogTag) Modify(id uint64, data request.BlogTagForm) error {
	tag := model.BlogTag{
		Name:        data.Name,
		Description: data.Description,
	}
	return global.DB.Model(model.BlogTag{}).Where("id = ?", id).Updates(&tag).Error
}

func (t BlogTag) Delete(id uint64) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("blog_tag_id = ?", id).Delete(&model.BlogPostTag{})
		return tx.Delete(&model.BlogTag{}, id).Error
	})
}

func (t BlogTag) Single(id uint64) *response.BlogTagResult {
	var tag model.BlogTag
	result := global.DB.Model(&model.BlogTag{}).Where("id = ?", id).First(&tag)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.LOGGER.Info("tag not found#", zap.Uint64("id", id))
		return nil
	}
	res := t.tagResponseData(tag)
	return &res
}

func (t BlogTag) List(query request.BlogTagQuery) (result []response.BlogTagResult, total int64) {
	tagModel := global.DB.Model(&model.BlogTag{})
	if query.Search != "" {
		tagModel.Where("name like ?", "%"+query.Search+"%")
	}
	err := tagModel.Count(&total).Error
	if err != nil {
		return
	}
	var tags []model.BlogTag
	tagModel = tagModel.Scopes(model.Paginate(query.Page, query.Limit)).Find(&tags)
	if tagModel.Error != nil {
		global.LOGGER.Error("tags find error#" + tagModel.Error.Error())
		return
	}
	for _, tag := range tags {
		postResult := t.tagResponseData(tag)
		result = append(result, postResult)
	}
	return
}

func (t BlogTag) tagResponseData(tag model.BlogTag) response.BlogTagResult {
	return response.BlogTagResult{
		ID:          tag.ID,
		Name:        tag.Name,
		Description: tag.Description,
	}
}
