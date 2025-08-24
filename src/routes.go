package src

import (
	"github.com/gin-gonic/gin"
	"github.com/zombox0633/printer_backend_go/src/mockApi"
)

func RoutersGroup(router *gin.Engine) {
	group := router.Group("/api")
	{
		mockApi.MockApiRoutes(group)
	}
}
