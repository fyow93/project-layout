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
	if err != nil {
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
