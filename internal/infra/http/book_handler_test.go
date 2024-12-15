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

	// Create a book to test retrieval via HTTP interface
	bookJSON := `{"id":"1", "title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Retrieve the book via HTTP interface
	req, _ = http.NewRequest("GET", "/book/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")

	// Delete the book via HTTP interface
	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}

func TestHandler_CreateBook(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	bookJSON := `{"id":"1", "title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "created")

	// Verify the book was created via HTTP interface
	req, _ = http.NewRequest("GET", "/book/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")

	// Delete the book via HTTP interface
	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}

func TestHandler_UpdateBook(t *testing.T) {
	handler := setupTestHandler(t)
	router := gin.Default()
	logger, _ := zap.NewProduction()
	handler.RegisterRoutes(router, logger)

	// Create a book to test update via HTTP interface
	bookJSON := `{"id":"1", "title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/book", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Update the book via HTTP interface
	updatedBookJSON := `{"title":"Updated Book", "author":"Updated Author"}`
	req, _ = http.NewRequest("PUT", "/book/1", strings.NewReader(updatedBookJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "updated")

	// Verify the book was updated via HTTP interface
	req, _ = http.NewRequest("GET", "/book/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Book")

	// Delete the book via HTTP interface
	req, _ = http.NewRequest("DELETE", "/book/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
}
