package blog

import (
	"net/http"

	"github.com/GanymedeNil/GoFrameworkBase/internal/request"
	"github.com/GanymedeNil/GoFrameworkBase/internal/service"
	"github.com/GanymedeNil/GoFrameworkBase/internal/util"

	"github.com/gin-gonic/gin"
)

//@Summary category list
//@Tags  admin/blog
//@Accept json
//@Param q query request.BlogCategoryQuery false "tags query"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse{result=util.Paginate{data=[]response.BlogCategoryResult}}
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/categories [get]
func CategoryList(c *gin.Context) {
	var query request.BlogCategoryQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	list, total := new(service.BlogCategory).List(query)
	result := util.Paginate{
		Total:       total,
		PerPage:     query.Limit,
		CurrentPage: query.Page,
		Data:        list,
	}
	util.NewResponse(c, http.StatusOK, "OK", result)
}

//@Summary create category
//@Tags  admin/blog
//@Accept json
//@Param data body request.BlogCategoryForm true "Body"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/categories [post]
func CategoryCreate(c *gin.Context) {
	var data request.BlogCategoryForm
	if err := c.ShouldBindJSON(&data); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogCategory).Create(data); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary modify category
//@Tags  admin/blog
//@Accept json
//@Param id path integer true "BlogCategory ID"
//@Param data body request.BlogCategoryForm true "Body"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/categories/{id} [put]
func CategoryModify(c *gin.Context) {
	var data request.BlogCategoryForm
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogCategory).Modify(uriID.ID, data); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary delete categories
//@Tags  admin/blog
//@Accept json
//@Param id path integer true "BlogCategories ID"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/categories/{id} [delete]
func CategoryDelete(c *gin.Context) {
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogCategory).Delete(uriID.ID); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary get single category
//@Tags admin/blog
//@Accept json
//@Param id path integer true "BlogCategory ID"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse{result=response.BlogCategoryResult}
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/categories/{id} [get]
func CategorySingle(c *gin.Context) {
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	result := new(service.BlogCategory).Single(uriID.ID)
	util.NewResponse(c, http.StatusOK, "OK", result)
}
