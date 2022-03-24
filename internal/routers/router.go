package routers

import (
	"ArtistHome/internal/routers/admin"
	"ArtistHome/internal/routers/web"

	"github.com/gin-gonic/gin"
)

func Create(engine *gin.Engine) {
	web.Routes(engine)
	admin.Routes(engine)
}
