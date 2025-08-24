package mockApi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zombox0633/printer_backend_go/src/utils"
)

type MockApiControllersType struct {
	s *MockApiServicesType
}

func NewMockApiControllers(s *MockApiServicesType) *MockApiControllersType {
	return &MockApiControllersType{s: s}
}

func (c *MockApiControllersType) GetAllMockApi(ctx *gin.Context) {
	item, err := c.s.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items 😿"})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c *MockApiControllersType) GetByIDMockApi(ctx *gin.Context) {
	id, ok := utils.ParamInt64(ctx, "id")
	if !ok {
		return
	}

	item, err := c.s.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID %d not found 😿", id)})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c *MockApiControllersType) CreateMockApi(ctx *gin.Context) {
	var req MockApiRequestType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data 😿 : " + err.Error()})
		return
	}

	item, err := c.s.Create(req.Title)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

func (c *MockApiControllersType) UpdateMockApi(ctx *gin.Context) {
	id, ok := utils.ParamInt64(ctx, "id")
	if !ok {
		return
	}

	var req MockApiRequestType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data 😿 : " + err.Error()})
		return
	}

	item, err := c.s.Update(id, req.Title)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID %d not found 😿", id)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (c *MockApiControllersType) SoftDeleteMockApi(ctx *gin.Context) {
	id, ok := utils.ParamInt64(ctx, "id")
	if !ok {
		return
	}
	if err := c.s.SoftDelete(id); err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID %d not found 😿", id)})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Item with ID %d deleted successfully 😸", id),
	})
}
