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

func main() {
	logger, _ := middleware.NewLogger()
	defer logger.Sync()

	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	userRepo := repository.NewUserRepositoryImpl(db)
	bookRepo := repository.NewBookRepositoryImpl(db)

	err = userRepo.Initialize()
	if err != nil {
		logger.Fatal("Failed to create user table", zap.Error(err))
	}

	err = bookRepo.Initialize()
	if err != nil {
		logger.Fatal("Failed to create book table", zap.Error(err))
	}

	userDomainService := dservice.NewUserService(userRepo)
	userUseCaseFactory := usecases.NewUserUseCaseFactory(userDomainService)
	serviceFactory := service.NewServiceFactory(nil, userUseCaseFactory)
	userService := serviceFactory.CreateUserService()
	userHandler := http.NewUserHandler(userService)

	bookDomainService := dservice.NewBookService(bookRepo)
	bookUseCaseFactory := usecases.NewBookUseCaseFactory(bookDomainService)
	serviceFactory = service.NewServiceFactory(bookUseCaseFactory, nil)
	bookService := serviceFactory.CreateBookService()
	bookHandler := http.NewBookHandler(bookService)

	router := gin.Default()
	userHandler.RegisterRoutes(router, logger)
	bookHandler.RegisterRoutes(router, logger)

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
