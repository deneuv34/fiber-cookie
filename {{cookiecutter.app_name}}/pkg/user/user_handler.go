package user

import (
	"github.com/gofiber/fiber"
	"net/http"
)

type IUserHandler interface {
	ListUser(c *fiber.Ctx)
}

type Handler struct {
	service IUserService
}

func NewHandler(service IUserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ListUser(c *fiber.Ctx) {
	user, _ := h.service.FetchUser(c)
	c.Send(http.StatusOK, user)
}
