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
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// BookHandler 处理与书籍相关的 HTTP 请求
// 它作为应用程序和外部世界之间的接口层
type BookHandler struct {
	appService *service.BookService // 应用服务层引用
}

// NewBookHandler 创建新的书籍处理器
// 参数:
//   - appService: 书籍应用服务实例
//
// 返回:
//   - *BookHandler: 新创建的书籍处理器
func NewBookHandler(appService *service.BookService) *BookHandler {
	return &BookHandler{appService: appService}
}

// RegisterRoutes 注册所有书籍相关的路由到 Gin 路由器
// 参数:
//   - router: Gin 路由引擎
//   - logger: 日志记录器
func (h *BookHandler) RegisterRoutes(router *gin.Engine, logger *zap.Logger) {
	// 添加全局中间件
	router.Use(middleware.CORSMiddleware())         // 跨域资源共享中间件
	router.Use(middleware.LoggerMiddleware(logger)) // 日志中间件

	// 注册书籍相关的 RESTful API 路由
	router.GET("/book/:id", h.GetBook)       // 获取单本书籍
	router.POST("/book", h.CreateBook)       // 创建新书籍
	router.PUT("/book/:id", h.UpdateBook)    // 更新书籍
	router.DELETE("/book/:id", h.DeleteBook) // 删除书籍
}

// GetBook 处理获取单本书籍的 HTTP 请求
// 对应 GET /book/:id 路由
// 参数:
//   - c: Gin 上下文
func (h *BookHandler) GetBook(c *gin.Context) {
	// 从 URL 路径中获取书籍 ID
	id := c.Param("id")

	// 调用应用服务层查找书籍
	book, err := h.appService.ExecuteFindByID(id)
	if err != nil {
		// 如果发生错误，返回 500 状态码和错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 成功找到书籍，返回 200 状态码和书籍数据
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.appService.ExecuteCreate(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id
	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.appService.ExecuteUpdate(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.appService.ExecuteDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
