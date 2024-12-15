// Package main 包含应用程序的入口点。
// 这是应用程序的启动层，负责初始化各个层次并启动 HTTP 服务器。
// 上一层：无（这是最外层）
// 下一层：接口层

package main

import (
	"project-layout/internal/application/service"
	"project-layout/internal/application/usecases"
	dservice "project-layout/internal/domain/service"
	"project-layout/internal/infra/http"
	"project-layout/internal/infra/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	repo, err := repository.NewRepositoryImpl("file::memory:?cache=shared")
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	err = repo.Initialize()
	if err != nil {
		logger.Fatal("Failed to create table", zap.Error(err))
	}

	domainService := dservice.NewDomainService(repo)
	useCaseFactory := usecases.NewUseCaseFactory(domainService)
	serviceFactory := service.NewServiceFactory(useCaseFactory)
	appService := serviceFactory.CreateApplicationService()
	handler := http.NewHandler(appService)

	router := gin.Default()
	handler.RegisterRoutes(router, logger)

	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
