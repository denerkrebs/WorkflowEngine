package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/http/handler"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Rotas
	r.Get("/health", healthHandler)

	r.Mount("/users", userRoutes())

	return r
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func userRoutes() chi.Router {
	r := chi.NewRouter()

    userHandler := handler.NewUserHandler();

	r.Post("/", userHandler.NewUser)

	return r
}
