package service

import (
	"errors"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"
	"github.com/GanymedeNil/GoFrameworkBase/internal/model"
	"github.com/GanymedeNil/GoFrameworkBase/internal/request"
	"github.com/GanymedeNil/GoFrameworkBase/internal/response"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BlogCategory struct {
}

func (c BlogCategory) Create(data request.BlogCategoryForm) error {
	category := model.BlogCategory{
		Name:        data.Name,
		Description: data.Description,
	}
	return global.DB.Create(&category).Error

}

func (c BlogCategory) Modify(id uint64, data request.BlogCategoryForm) error {
	category := model.BlogCategory{
		Name:        data.Name,
		Description: data.Description,
	}
	return global.DB.Model(model.BlogCategory{}).Where("id = ?", id).Updates(&category).Error
}

func (c BlogCategory) Delete(id uint64) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("blog_category_id = ?", id).Delete(&model.BlogPostCategory{})
		return tx.Delete(&model.BlogCategory{}, id).Error
	})
}

func (c BlogCategory) Single(id uint64) *response.BlogCategoryResult {
	var category model.BlogCategory
	result := global.DB.Model(&model.BlogCategory{}).Where("id = ?", id).First(&category)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		global.LOGGER.Info("category not found#", zap.Uint64("id", id))
		return nil
	}
	res := c.categoryResponseData(category)
	return &res
}

func (c BlogCategory) List(query request.BlogCategoryQuery) (result []response.BlogCategoryResult, total int64) {
	categoryModel := global.DB.Model(&model.BlogCategory{})
	if query.Search != "" {
		categoryModel.Where("name like ?", "%"+query.Search+"%")
	}
	err := categoryModel.Count(&total).Error
	if err != nil {
		return
	}
	var categories []model.BlogCategory
	categoryModel = categoryModel.Scopes(model.Paginate(query.Page, query.Limit)).Find(&categories)
	if categoryModel.Error != nil {
		global.LOGGER.Error("categories find error#" + categoryModel.Error.Error())
		return
	}
	for _, category := range categories {
		postResult := c.categoryResponseData(category)
		result = append(result, postResult)
	}
	return
}

func (c BlogCategory) categoryResponseData(category model.BlogCategory) response.BlogCategoryResult {
	return response.BlogCategoryResult{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
