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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func setupTestHandler(t *testing.T) *BookHandler {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	repo := repository.NewBookRepositoryImpl(db)
	err = repo.Initialize()
	if err != nil {
		color.Red("Failed to create table: %v", err)
		t.Fatalf("Failed to create table: %v", err)
	}

	// Register cleanup function
	t.Cleanup(func() {
		repo.DeleteAllBooks()
		repo.Shutdown()
	})

	bookService := dservice.NewBookService(repo)
	bookUseCase := usecases.NewBookUseCase(bookService)
	appService := service.NewBookService(bookUseCase)
	return NewBookHandler(appService)
}

func TestHandler_GetBook(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// 使用有效的 UUID v4
	validUUID := uuid.New().String()
	bookJSON := `{"id":"` + validUUID + `", "title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())

	// Retrieve the book via HTTP interface
	req, _ = http.NewRequest("GET", "/book/"+validUUID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())
	assert.Contains(t, w.Body.String(), "Test Book", "Response: %s", w.Body.String())

	// Delete the book via HTTP interface
	req, _ = http.NewRequest("DELETE", "/book/"+validUUID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())
}

func TestHandler_CreateBook(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// 使用有效的 UUID v4
	validUUID := uuid.New().String()
	bookJSON := `{"id":"` + validUUID + `", "title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())
	assert.Contains(t, w.Body.String(), "created", "Response: %s", w.Body.String())

	// Verify the book was created via HTTP interface
	req, _ = http.NewRequest("GET", "/book/"+validUUID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())
	assert.Contains(t, w.Body.String(), "Test Book", "Response: %s", w.Body.String())

	// Delete the book via HTTP interface
	req, _ = http.NewRequest("DELETE", "/book/"+validUUID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())
	assert.Contains(t, w.Body.String(), "deleted", "Response: %s", w.Body.String())
}

func TestHandler_UpdateBook(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// 使用有效的 UUID v4
	validUUID := uuid.New().String()
	bookJSON := `{"id":"` + validUUID + `", "title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "Response: %s", w.Body.String())

	// Update the book with invalid data
	invalidBookJSON := `{"id":"` + validUUID + `", "title":"", "author":""}`
	req, _ = http.NewRequest("PUT", "/book/"+validUUID, strings.NewReader(invalidBookJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "Response: %s", w.Body.String())
}
