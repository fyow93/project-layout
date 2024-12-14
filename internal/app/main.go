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

	router.Run(":8080")
}
