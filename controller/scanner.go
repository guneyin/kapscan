package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/logger"
	"github.com/guneyin/kapscan/repo/company"
	"github.com/guneyin/kapscan/service/scanner"
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

	grp.Get("/sync-symbols", s.SyncSymbols)
	grp.Get("/sync-company", s.SyncCompany)

	return s
}

func (s *Scanner) SyncSymbols(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 0)

	go func(ctx context.Context) {
		err := s.svc.SyncSymbolList(c.Context(), limit)
		if err != nil {
			logger.Log().ErrorContext(ctx, err.Error())
		}
	}(c.Context())

	c.Append("HX-Refresh", "true")
	return c.JSON(fiber.Map{"status": "ok", "message": "sync started"})
}

func (s *Scanner) SyncCompany(c *fiber.Ctx) error {
	code := c.Query("code")

	companyRepo := company.NewRepo()
	cmp, err := companyRepo.GetByCode(c.Context(), code)
	if err != nil {
		return err
	}

	err = s.svc.SyncCompanyWithShares(c.Context(), cmp)
	if err != nil {
		return err
	}

	c.Append("HX-Refresh", "true")
	return c.JSON(fiber.Map{"status": "ok", "message": "sync started"})
}
