package core

import (
	"ArtistHome/internal/global"
	"ArtistHome/internal/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Service() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization", " x-requested-with")
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	routers.Create(router)
	router.Run(global.CONFIG.App.Addr)
}
