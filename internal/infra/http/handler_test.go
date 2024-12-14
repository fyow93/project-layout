package http

import (
	"net/http"
	"net/http/httptest"
	"project-layout/internal/application/service"
	"project-layout/internal/application/usecases"
	dservice "project-layout/internal/domain/service"
	"project-layout/internal/infra/repository"
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func setupTestHandler(t *testing.T) *Handler {
	db, err := repository.NewRepositoryImpl("file::memory:?cache=shared")
	if err != nil {
		color.Red("Failed to initialize database: %v", err)
		t.Fatalf("Failed to initialize database: %v", err)
	}
	err = db.Initialize()
	if err != nil {
		color.Red("Failed to create table: %v", err)
		t.Fatalf("Failed to create table: %v", err)
	}

	// Register cleanup function
	t.Cleanup(func() {
		db.DeleteAllEntities()
		db.Shutdown()
	})

	domainService := dservice.NewDomainService(db)
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

	// Create an entity to test retrieval via HTTP interface
	entityJSON := `{"id":"1", "name":"Test Entity"}`
	req, _ := http.NewRequest("POST", "/entity", strings.NewReader(entityJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Retrieve the entity via HTTP interface
	req, _ = http.NewRequest("GET", "/entity/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Entity")

	// Delete the entity via HTTP interface
	req, _ = http.NewRequest("DELETE", "/entity/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
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

	// Verify the entity was created via HTTP interface
	req, _ = http.NewRequest("GET", "/entity/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Entity")

	// Delete the entity via HTTP interface
	req, _ = http.NewRequest("DELETE", "/entity/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}

func TestHandler_UpdateEntity(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// Create an entity to test update via HTTP interface
	entityJSON := `{"id":"1", "name":"Test Entity"}`
	req, _ := http.NewRequest("POST", "/entity", strings.NewReader(entityJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Update the entity via HTTP interface
	updatedEntityJSON := `{"name":"Updated Entity"}`
	req, _ = http.NewRequest("PUT", "/entity/1", strings.NewReader(updatedEntityJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "updated")

	// Verify the entity was updated via HTTP interface
	req, _ = http.NewRequest("GET", "/entity/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Entity")

	// Delete the entity via HTTP interface
	req, _ = http.NewRequest("DELETE", "/entity/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}
