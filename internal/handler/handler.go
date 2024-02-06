package handler

import (
	"test-crud-api/internal/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	Services service.Service
}

func NewHandler(services service.Service) *Handler {
	return &Handler{Services: services}
}
func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.createUser)
	r.Get("/users", h.getAllUsersWithFilters)
	r.Get("/users/", h.findAllUsers)
	r.Get("/{id}", h.getUserByID)
	r.Post("/users/{id}", h.deleteUser)

	return r
}
