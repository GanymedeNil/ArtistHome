package admin

import (
	"ArtistHome/internal/middleware"
	"ArtistHome/internal/routers/admin/blog"
	"ArtistHome/internal/routers/admin/user"
	"ArtistHome/internal/util"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	admin := router.Group("/admin")
	admin.POST("/login", middleware.Jwt().LoginHandler)
	admin.Use(middleware.Jwt().MiddlewareFunc())
	admin.GET("/me", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		util.NewResponse(c, http.StatusOK, "OK", claims)
	})
	admin.GET("/getAsyncRoutes", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0, "info": []map[string]interface{}{},
		})
	})
	admin.GET("/logout", middleware.Jwt().LogoutHandler)
	admin.POST("/refresh_token", middleware.Jwt().RefreshHandler)
	blog.Routes(admin)
	user.Routes(admin)
}
