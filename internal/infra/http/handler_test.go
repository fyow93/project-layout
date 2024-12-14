package http

import (
	"net/http"
	"net/http/httptest"
	"project-layout/internal/application/service"
	"project-layout/internal/application/usecases"
	"project-layout/internal/domain/model"
	dservice "project-layout/internal/domain/service"
	"project-layout/internal/infra/repository"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func setupTestHandler(t *testing.T) *Handler {
	repo, err := repository.NewRepositoryImpl("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	err = repo.Initialize()
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}
	domainService := dservice.NewDomainService(repo)
	useCaseFactory := usecases.NewUseCaseFactory(domainService)
	serviceFactory := service.NewServiceFactory(useCaseFactory)
	appService := serviceFactory.CreateApplicationService()

	return NewHandler(appService)
}

func TestHandler_GetEntity(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// Create an entity to test retrieval
	entity := model.Entity{ID: "1", Name: "Test Entity"}
	handler.appService.ExecuteCreate(entity)

	req, _ := http.NewRequest("GET", "/entity/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Entity")
}

func TestHandler_CreateEntity(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	entityJSON := `{"id":"1", "name":"Test Entity"}`
	req, _ := http.NewRequest("POST", "/entity", strings.NewReader(entityJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "created")

	// Verify the entity was created
	savedEntity, err := handler.appService.ExecuteFindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, "Test Entity", savedEntity.Name)
}
