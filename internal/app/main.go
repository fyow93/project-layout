// Package main 包含应用程序的入口点。
// 这是应用程序的启动层，负责初始化各个层次并启动 HTTP 服务器。
// 上一层：无（这是最外层）
// 下一层：接口层

package main

import (
	"database/sql"
	"project-layout/internal/application/service"
	"project-layout/internal/application/usecases"
	dservice "project-layout/internal/domain/service"
	"project-layout/internal/infra/http"
	"project-layout/internal/infra/http/middleware"
	"project-layout/internal/infra/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ServiceFactory 通过依赖注入创建和组织各个服务组件
// 该工厂遵循依赖注入原则，集中管理系统组件的创建和依赖关系
// 这种设计有助于:
// - 解耦系统组件
// - 简化单元测试
// - 提高代码的可维护性和可扩展性
type ServiceFactory struct {
	db     *sql.DB   // 数据库连接
	logger *zap.Logger // 日志记录器
}

// NewServiceFactory 创建一个新的服务工厂实例
// 参数:
//   - db: 数据库连接
//   - logger: 日志记录器
// 返回:
//   - *ServiceFactory: 新创建的服务工厂实例
func NewServiceFactory(db *sql.DB, logger *zap.Logger) *ServiceFactory {
	return &ServiceFactory{db: db, logger: logger}
}

// CreateUserHandler 创建并初始化用户处理器
// 该方法遵循 DDD 的依赖注入和分层架构原则，按照以下步骤构建组件:
// 1. 创建仓储实现
// 2. 初始化仓储
// 3. 创建领域服务
// 4. 创建用例工厂
// 5. 创建应用服务
// 6. 创建 HTTP 处理器
// 返回:
//   - *http.UserHandler: 用户 HTTP 处理器
//   - error: 如果初始化过程中发生错误
func (f *ServiceFactory) CreateUserHandler() (*http.UserHandler, error) {
	userRepo := repository.NewUserRepositoryImpl(f.db)
	if err := userRepo.Initialize(); err != nil {
		return &http.UserHandler{}, err
	}
	userDomainService := dservice.NewUserService(userRepo)
	userUseCaseFactory := usecases.NewUserUseCaseFactory(userDomainService)
	serviceFactory := service.NewServiceFactory(nil, userUseCaseFactory)
	userService := serviceFactory.CreateUserService()
	return http.NewUserHandler(userService), nil
}

// CreateBookHandler 创建并初始化书籍处理器
// 该方法遵循与 CreateUserHandler 相同的依赖注入和分层架构原则
// 返回:
//   - *http.BookHandler: 书籍 HTTP 处理器
//   - error: 如果初始化过程中发生错误
func (f *ServiceFactory) CreateBookHandler() (*http.BookHandler, error) {
	bookRepo := repository.NewBookRepositoryImpl(f.db)
	if err := bookRepo.Initialize(); err != nil {
		return &http.BookHandler{}, err
	}
	bookDomainService := dservice.NewBookService(bookRepo)
	bookUseCaseFactory := usecases.NewBookUseCaseFactory(bookDomainService)
	serviceFactory := service.NewServiceFactory(bookUseCaseFactory, nil)
	bookService := serviceFactory.CreateBookService()
	return http.NewBookHandler(bookService), nil
}

func main() {
	// 初始化日志系统
	logger, _ := middleware.NewLogger()
	defer logger.Sync() // 确保日志缓冲区在程序退出前被刷新

	// 初始化内存数据库连接
	// 注意：在生产环境中，通常应该使用持久化数据库
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	// 创建服务工厂，用于依赖注入
	factory := NewServiceFactory(db, logger)

	// 初始化用户处理器
	userHandler, err := factory.CreateUserHandler()
	if err != nil {
		logger.Fatal("Failed to initialize user handler", zap.Error(err))
	}

	// 初始化图书处理器
	bookHandler, err := factory.CreateBookHandler()
	if err != nil {
		logger.Fatal("Failed to initialize book handler", zap.Error(err))
	}

	// 设置 Gin Web 框架路由
	router := gin.Default()
	
	// 注册路由并关联处理函数
	userHandler.RegisterRoutes(router, logger)
	bookHandler.RegisterRoutes(router, logger)

	// 启动 HTTP 服务器在端口 8080
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
