package web

import (
	"github.com/GanymedeNil/GoFrameworkBase/internal/routers/web/blog"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	web := router.Group("/web")
	blog.Routes(web)
}
