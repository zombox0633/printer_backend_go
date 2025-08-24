package mockApi

import "github.com/gin-gonic/gin"

func MockApiRoutes(router *gin.RouterGroup) {
	mockApiRepo := NewMockRepository()
	mockApiService := NewMockApiServices(mockApiRepo)
	mockApiController := NewMockApiControllers(mockApiService)

	mockApiGroup := router.Group("/mock")
	{
		mockApiGroup.GET("", mockApiController.GetAllMockApi)
		mockApiGroup.GET("/:id", mockApiController.GetByIDMockApi)
		mockApiGroup.POST("", mockApiController.CreateMockApi)
		mockApiGroup.PUT("/:id", mockApiController.UpdateMockApi)
		mockApiGroup.POST("/delete/:id", mockApiController.SoftDeleteMockApi)

	}
}
