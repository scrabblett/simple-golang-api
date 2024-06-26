package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	mw "simple-golang-api/internal/handler/middleware"
	v1 "simple-golang-api/internal/handler/v1"
	"simple-golang-api/internal/service"
	"time"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(duration time.Duration) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.URLFormat,
		mw.ReqLogger(zap.L()),
		mw.TimeoutMiddleware(duration),
	)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *chi.Mux) {
	handlerV1 := v1.NewHandler(h.services)

	router.Route("/api", func(router chi.Router) {
		handlerV1.Init(router)
	})
}
