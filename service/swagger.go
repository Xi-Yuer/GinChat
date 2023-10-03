package service

import (
	"GinChat/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func Swagger() gin.HandlerFunc {
	// export PATH=$(go env GOPATH)/bin:$PATH
	docs.SwaggerInfo.BasePath = ""
	url := ginSwagger.URL("http://localhost:8080/utils/swagger/doc.json") // The url pointing to API definition
	return ginSwagger.WrapHandler(swaggerFiles.Handler, url)
}
