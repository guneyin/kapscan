package controller

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IController interface {
	name() string
	setRoutes(router fiber.Router) IController
}

type Controller struct {
	db          *gorm.DB
	router      fiber.Router
	controllers map[string]IController
}

func NewController(db *gorm.DB, router fiber.Router) *Controller {
	c := &Controller{
		db:          db,
		router:      router,
		controllers: make(map[string]IController),
	}
	c.registerControllers()

	return c
}

func (c Controller) registerControllers() {
	c.register(newGeneralController)
	c.register(newScannerController)
}

func (c Controller) register(f func(db *gorm.DB) IController) {
	hnd := f(c.db).setRoutes(c.router)
	c.controllers[hnd.name()] = hnd
}
