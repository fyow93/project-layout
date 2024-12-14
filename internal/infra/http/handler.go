// Package http 包含应用程序的 HTTP 处理程序。
// 这是接口层，负责处理 HTTP 请求和响应。
// 它与应用服务层交互以执行业务逻辑操作。
// 上一层：无（这是最外层）
// 下一层：应用服务层

// 在 DDD 中，接口层负责与外部世界（如用户界面、API 客户端等）进行交互。
// 它将外部请求转换为应用服务层的调用，并将应用服务层的结果返回给外部。
// 该层的核心定义是控制器（Controller），在这里体现为 HTTP 处理程序（Handler）。
// ***这是防腐层的一部分，用于隔离外部接口的变化。***

// Diagram:
// 外部世界（用户界面、API 客户端等） -> 接口层（HTTP 处理程序） -> 应用服务层

package http

import (
	"net/http"
	"project-layout/internal/application/service"
	"project-layout/internal/domain/model"
	"project-layout/internal/infra/http/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	appService *service.ApplicationService
}

func NewHandler(appService *service.ApplicationService) *Handler {
	return &Handler{appService: appService}
}

func (h *Handler) RegisterRoutes(router *gin.Engine, logger *zap.Logger) {
	router.Use(middleware.CORSMiddleware())
	//router.Use(middleware.AuthMiddleware())
	router.Use(middleware.LoggerMiddleware(logger))

	router.GET("/entity/:id", h.GetEntity)
	router.POST("/entity", h.CreateEntity)
	router.PUT("/entity/:id", h.UpdateEntity)
	router.DELETE("/entity/:id", h.DeleteEntity)
}

func (h *Handler) GetEntity(c *gin.Context) {
	id := c.Param("id")
	entity, err := h.appService.ExecuteFindByID(id)
	if (err != nil) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (h *Handler) CreateEntity(c *gin.Context) {
	var entity model.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.appService.ExecuteCreate(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

func (h *Handler) UpdateEntity(c *gin.Context) {
	id := c.Param("id")
	var entity model.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity.ID = id
	if err := h.appService.ExecuteUpdate(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteEntity(c *gin.Context) {
	id := c.Param("id")
	if err := h.appService.ExecuteDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
