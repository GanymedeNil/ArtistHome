package blog

import (
	"ArtistHome/internal/request"
	"ArtistHome/internal/service"
	"ArtistHome/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//@Summary tag list
//@Tags  admin/blog
//@Accept json
//@Param q query request.BlogTagQuery false "tags query"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse{result=util.Paginate{data=[]response.BlogTagResult}}
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/tags [get]
func TagList(c *gin.Context) {
	var query request.BlogTagQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	list, total := new(service.BlogTag).List(query)
	result := util.Paginate{
		Total:       total,
		PerPage:     query.Limit,
		CurrentPage: query.Page,
		Data:        list,
	}
	util.NewResponse(c, http.StatusOK, "OK", result)
}

//@Summary create tag
//@Tags  admin/blog
//@Accept json
//@Param data body request.BlogTagForm true "Body"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/tags [post]
func TagCreate(c *gin.Context) {
	var data request.BlogTagForm
	if err := c.ShouldBindJSON(&data); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogTag).Create(data); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary modify tag
//@Tags  admin/blog
//@Accept json
//@Param id path integer true "BlogTags ID"
//@Param data body request.BlogTagForm true "Body"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/tags/{id} [put]
func TagModify(c *gin.Context) {
	var data request.BlogTagForm
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogTag).Modify(uriID.ID, data); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary delete tag
//@Tags  admin/blog
//@Accept json
//@Param id path integer true "BlogTag ID"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/tags/{id} [delete]
func TagDelete(c *gin.Context) {
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogTag).Delete(uriID.ID); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary get single tag
//@Tags admin/blog
//@Accept json
//@Param id path integer true "BlogTag ID"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse{result=response.BlogTagResult}
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/tags/{id} [get]
func TagSingle(c *gin.Context) {
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	result := new(service.BlogTag).Single(uriID.ID)
	util.NewResponse(c, http.StatusOK, "OK", result)
}
