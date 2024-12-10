package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/mw"
	"github.com/guneyin/kapscan/internal/service/scanner"
	"gorm.io/gorm"
)

const scannerControllerName = "scanner"

type Scanner struct {
	svc *scanner.Service
}

func newScannerController(db *gorm.DB) IController {
	svc := scanner.NewScannerService(db)

	return &Scanner{svc}
}

func (s *Scanner) name() string {
	return scannerControllerName
}

func (s *Scanner) setRoutes(router fiber.Router) IController {
	grp := router.Group(s.name())
	grp.Get("/:symbol", s.Scan)
	return s
}

func (s *Scanner) Scan(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	res, err := s.svc.Scan(c.Context(), symbol)
	if err != nil {
		return mw.Error(c, err)
	}

	return mw.OK(c, "symbol fetched", res)
}
