package core

import (
	"github.com/GanymedeNil/GoFrameworkBase/internal/global"
	"github.com/GanymedeNil/GoFrameworkBase/internal/middleware"
	"github.com/GanymedeNil/GoFrameworkBase/internal/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Service() {
	config := cors.DefaultConfig()
	if global.CONFIG.App.Debug && global.CONFIG.App.Env == global.DevMode {
		gin.SetMode(gin.DebugMode)
		config.AllowAllOrigins = true
	} else {
		if global.CONFIG.App.Env == global.TestMode {
			gin.SetMode(gin.TestMode)
			config.AllowAllOrigins = true
		} else if global.CONFIG.App.Env == global.ReleaseMode {
			gin.SetMode(gin.ReleaseMode)
		}
	}
	router := gin.Default()
	config.AddAllowHeaders("Authorization", " x-requested-with")
	router.Use(cors.New(config))
	router.Use(middleware.Logger(), middleware.Recovery(true))
	routers.Create(router)
	router.Run(global.CONFIG.App.Addr)
}
