package controller

import (
	"github.com/gofiber/fiber/v2"
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
	return s
}
