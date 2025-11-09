package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/container"
	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/http/handler"
)

func NewRouter(c *container.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Get("/health", healthHandler)

	r.Mount("/users", userRoutes(c))

	return r
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func userRoutes(c *container.Container) chi.Router {
	r := chi.NewRouter()

	userHandler := handler.NewUserHandler(*c.RegisterUserUseCase)

	r.Post("/", userHandler.NewUser)

	return r
}
