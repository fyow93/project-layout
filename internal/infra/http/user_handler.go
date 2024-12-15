// Package http 包含应用程序的 HTTP 处理程序。
// 这是接口层，负责处理 HTTP 请求和响应。
// 它与应用服务层交互以执行业务逻辑操作。
// 上一层：无（这是最外层）
// 下一层：应用服务层

package http

import (
	"net/http"
	"project-layout/internal/application/service"
	"project-layout/internal/domain/model"
	"project-layout/internal/infra/http/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	appService *service.UserService
}

func NewUserHandler(appService *service.UserService) *UserHandler {
	return &UserHandler{appService: appService}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine, logger *zap.Logger) {
	router.Use(middleware.CORSMiddleware())
	//router.Use(middleware.AuthMiddleware())
	router.Use(middleware.LoggerMiddleware(logger))

	router.GET("/user/:id", h.GetUser)
	router.POST("/user", h.CreateUser)
	router.PUT("/user/:id", h.UpdateUser)
	router.DELETE("/user/:id", h.DeleteUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.appService.ExecuteFindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.appService.ExecuteCreate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := h.appService.ExecuteUpdate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.appService.ExecuteDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
