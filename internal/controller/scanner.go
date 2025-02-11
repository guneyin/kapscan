package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/service/scanner"
)

const controllerName = "scanner"

type Scanner struct {
	svc *scanner.Service
}

func newScannerController() IController {
	svc := scanner.NewService()

	return &Scanner{svc}
}

func (s *Scanner) name() string {
	return controllerName
}

func (s *Scanner) setRoutes(router fiber.Router) IController {
	grp := router.Group(s.name())

	grp.Get("/sync", s.Sync)

	return s
}

func (s *Scanner) Sync(c *fiber.Ctx) error {
	go func(ctx context.Context) {
		_ = s.svc.SyncCompanyList(ctx)
	}(c.Context())

	return c.JSON(fiber.Map{"status": "ok", "message": "sync started"})
}
