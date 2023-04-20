package v1

import (
	"avito_test_GO/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine,
	u usecase.UserContract,
	r usecase.ReportContract) {

	h := handler.Group("/v1")

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{
		newUserRoutes(h, u, r)
	}
}
