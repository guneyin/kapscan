package handler

import (
	"github.com/gofiber/fiber/v2"
)

type IHandler interface {
	name() string
	setRoutes(router fiber.Router) IHandler
}

type Handler struct {
	router      fiber.Router
	controllers map[string]IHandler
}

func NewWebHandler(router fiber.Router) *Handler {
	c := &Handler{
		router:      router,
		controllers: make(map[string]IHandler),
	}
	c.registerWebHandlers()

	return c
}

func (c Handler) registerWebHandlers() {
	c.register(newCompanyWebHandler)
}

func (c Handler) register(f func() IHandler) {
	hnd := f().setRoutes(c.router)
	c.controllers[hnd.name()] = hnd
}
