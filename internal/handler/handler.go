package handler

import (
	"test-crud-api/internal/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.createUser)
	r.Get("/users", h.getAllUsersWithFilters)
	r.Get("/users/all", h.findAllUsers)
	r.Get("/users/{id}", h.getUserByID)
	r.Post("/users/{id}", h.deleteUser)

	return r
}
