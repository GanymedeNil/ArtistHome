package main

import (
	_ "ArtistHome/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json"))
	ginSwagger.DefaultModelsExpandDepth(-1)
	//api.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8081")
}
