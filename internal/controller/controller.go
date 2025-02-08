package controller

import (
	"github.com/gofiber/fiber/v2"
)

type IController interface {
	name() string
	setRoutes(router fiber.Router) IController
}

type Controller struct {
	router      fiber.Router
	controllers map[string]IController
}

func NewController(router fiber.Router) *Controller {
	c := &Controller{
		router:      router,
		controllers: make(map[string]IController),
	}
	c.registerControllers()

	return c
}

func (c Controller) registerControllers() {
	c.register(newGeneralController)
	c.register(newCompanyController)
}

func (c Controller) register(f func() IController) {
	hnd := f().setRoutes(c.router)
	c.controllers[hnd.name()] = hnd
}
