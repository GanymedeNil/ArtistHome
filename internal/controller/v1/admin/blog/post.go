package blog

import (
	"ArtistHome/internal/request"
	"ArtistHome/internal/service"
	"ArtistHome/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//@Summary post list
//@Tags  admin/blog
//@Accept json
//@Param q query request.BlogPostQuery false "posts query"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse{result=util.Paginate{data=[]response.BlogPostResult}}
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/posts [get]
func PostList(c *gin.Context) {
	var query request.BlogPostQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	list, total := new(service.BlogPost).List(query)
	result := util.Paginate{
		Total:       total,
		PerPage:     query.Limit,
		CurrentPage: query.Page,
		Data:        list,
	}
	util.NewResponse(c, http.StatusOK, "OK", result)
}

//@Summary create post
//@Tags  admin/blog
//@Accept json
//@Param data body request.BlogPostForm true "Body"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/posts [post]
func PostCreate(c *gin.Context) {
	var data request.BlogPostForm
	if err := c.ShouldBindJSON(&data); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogPost).Create(data); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary modify post
//@Tags  admin/blog
//@Accept json
//@Param id path integer true "BlogPost ID"
//@Param data body request.BlogPostForm true "Body"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/posts/{id} [put]
func PostModify(c *gin.Context) {
	var data request.BlogPostForm
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogPost).Modify(uriID.ID, data); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary delete post
//@Tags  admin/blog
//@Accept json
//@Param id path integer true "BlogPost ID"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/posts/{id} [delete]
func PostDelete(c *gin.Context) {
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := new(service.BlogPost).Delete(uriID.ID); err != nil {
		util.NewError(c, http.StatusInternalServerError, err)
		return
	}
	util.NewResponse(c, http.StatusOK, "OK", struct{}{})
}

//@Summary get single post
//@Tags admin/blog
//@Accept json
//@Param id path integer true "BlogPost ID"
//@Security bearerAuth
//@Produce json
//@Success 200 {object} util.JsonResponse{result=response.BlogPostResult}
//@Failure 400 {object} util.JsonResponse
//@Failure 500 {object} util.JsonResponse
//@Router /admin/blog/v1/posts/{id} [get]
func PostSingle(c *gin.Context) {
	var uriID request.UriID
	if err := c.ShouldBindUri(&uriID); err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}
	result := new(service.BlogPost).Single(uriID.ID)
	util.NewResponse(c, http.StatusOK, "OK", result)
}
