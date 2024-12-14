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

	repo := &repository.RepositoryImpl{}
	domainService := dservice.NewDomainService(repo)
	useCase := usecases.NewUseCase(domainService)
	appService := service.NewApplicationService(useCase)
	handler := http.NewHandler(appService)

	router := gin.Default()
	handler.RegisterRoutes(router, logger)

	router.Run(":8080")
}
