package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/mw"
	"github.com/guneyin/kapscan/service/general"
)

const generalControllerName = "general"

type General struct {
	svc *general.Service
}

func newGeneralController() IController {
	return &General{general.NewGeneralService()}
}

func (g General) name() string {
	return generalControllerName
}

func (g General) setRoutes(router fiber.Router) IController {
	gr := router.Group(g.name())
	gr.Get("status", g.GeneralStatus)

	return g
}

func (g General) GeneralStatus(c *fiber.Ctx) error {
	return mw.OK(c, "service status fetched", g.svc.Status())
}
