package http

import (
	"net/http"
	"project-layout/internal/application/service"
)

type Handler struct {
	appService *service.ApplicationService
}

func NewHandler(appService *service.ApplicationService) *Handler {
	return &Handler{appService: appService}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle HTTP request
}
