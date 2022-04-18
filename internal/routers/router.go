package routers

import (
	"github.com/GanymedeNil/GoFrameworkBase/internal/routers/admin"
	"github.com/GanymedeNil/GoFrameworkBase/internal/routers/web"

	"github.com/gin-gonic/gin"
)

func Create(engine *gin.Engine) {
	web.Routes(engine)
	admin.Routes(engine)
}
