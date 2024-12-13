package main

import (
	"project-layout/internal/application/service"
	"project-layout/internal/application/usecases"
	dservice "project-layout/internal/domain/service"
	"project-layout/internal/infra/http"
	"project-layout/internal/infra/repository"
)

func main() {
	repo := &repository.RepositoryImpl{}
	domainService := dservice.NewDomainService(repo)
	useCase := usecases.NewUseCase(domainService)
	appService := service.NewApplicationService(useCase)
	handler := http.NewHandler(appService)

	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
