package v1

import (
	"awesomeProject/internal/handler/middleware"
	"awesomeProject/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(router chi.Router) {
	router.Route("/v1", func(router chi.Router) {
		//public routes
		h.initAuthRoutes(router)

		// private routes
		router.Group(func(router chi.Router) {
			router.Use(middleware.JwtMiddleware())
			h.initBooksRoutes(router)
		})
	})
}