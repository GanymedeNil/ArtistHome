package web

import (
	"ArtistHome/internal/routers/web/blog"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	web := router.Group("/web")
	blog.Routes(web)
}
