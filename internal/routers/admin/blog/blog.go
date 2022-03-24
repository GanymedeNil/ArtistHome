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
	v1.POST("/posts", blog.PostCreate)
	v1.GET("/posts/:id", blog.PostSingle)
	v1.PUT("/posts/:id", blog.PostModify)
	v1.DELETE("/posts/:id", blog.PostDelete)

	v1.GET("/tags", blog.TagList)
	v1.POST("/tags", blog.TagCreate)
	v1.GET("/tags/:id", blog.TagSingle)
	v1.PUT("/tags/:id", blog.TagModify)
	v1.DELETE("/tags/:id", blog.TagDelete)

	v1.GET("/categories", blog.CategoryList)
	v1.POST("/categories", blog.CategoryCreate)
	v1.GET("/categories/:id", blog.CategorySingle)
	v1.PUT("/categories/:id", blog.CategoryModify)
	v1.DELETE("/categories/:id", blog.CategoryDelete)
}
