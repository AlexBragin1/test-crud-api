package handler

import (
	"test-crud-api/internal/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func newHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.createUser)
	r.Post("/{id}", h.deleteUser)
	//r.GET("/users", h.getAllUsersWithFilter)
	r.Get("/users/{id}", h.getUserByID)
	r.Post("/{id}", h.deleteUser)

	return r
}
