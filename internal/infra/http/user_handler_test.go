package http

import (
	"database/sql"
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

func setupTestUserHandler(t *testing.T) *UserHandler {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	repo := repository.NewUserRepositoryImpl(db)
	err = repo.Initialize()
	if err != nil {
		color.Red("Failed to create table: %v", err)
		t.Fatalf("Failed to create table: %v", err)
	}

	// Register cleanup function
	t.Cleanup(func() {
		repo.DeleteAllUsers()
		repo.Shutdown()
	})

	userService := dservice.NewUserService(repo)
	useCase := usecases.NewUserUseCase(userService)
	appService := service.NewUserService(useCase)

	return NewUserHandler(appService)
}

func TestHandler_GetUser(t *testing.T) {
	handler := setupTestUserHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// Create a user to test retrieval via HTTP interface
	userJSON := `{"id":"1", "name":"Test User", "email":"test@example.com"}`
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Retrieve the user via HTTP interface
	req, _ = http.NewRequest("GET", "/user/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test User")

	// Delete the user via HTTP interface
	req, _ = http.NewRequest("DELETE", "/user/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}

func TestHandler_CreateUser(t *testing.T) {
	handler := setupTestUserHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	userJSON := `{"id":"1", "name":"Test User", "email":"test@example.com"}`
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "created")

	// Verify the user was created via HTTP interface
	req, _ = http.NewRequest("GET", "/user/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test User")

	// Delete the user via HTTP interface
	req, _ = http.NewRequest("DELETE", "/user/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}

func TestHandler_UpdateUser(t *testing.T) {
	handler := setupTestUserHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// Create a user to test update via HTTP interface
	userJSON := `{"id":"1", "name":"Test User", "email":"test@example.com"}`
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Update the user via HTTP interface
	updatedUserJSON := `{"name":"Updated User", "email":"updated@example.com"}`
	req, _ = http.NewRequest("PUT", "/user/1", strings.NewReader(updatedUserJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "updated")

	// Verify the user was updated via HTTP interface
	req, _ = http.NewRequest("GET", "/user/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated User")

	// Delete the user via HTTP interface
	req, _ = http.NewRequest("DELETE", "/user/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}
