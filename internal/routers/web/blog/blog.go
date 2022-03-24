package blog

import (
	"ArtistHome/internal/controller/v1/admin/blog"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	blogGroup := router.Group("/blog")
	V1(blogGroup)
}

func V1(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	v1.GET("/posts", blog.PostList)
	v1.GET("/posts/:id", blog.PostSingle)
	v1.GET("/tags", blog.TagList)
	v1.GET("/tags/:id", blog.TagSingle)
	v1.GET("/categories", blog.CategoryList)
	v1.GET("/categories/:id", blog.CategorySingle)
}
