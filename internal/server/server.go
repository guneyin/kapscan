package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/guneyin/kapscan/internal/mw"
	"time"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func NewServer(appName string) *fiber.App {
	app := fiber.New(fiber.Config{
		ServerHeader:      fmt.Sprintf("%s HTTP Server", appName),
		BodyLimit:         16 * 1024 * 1024,
		AppName:           appName,
		EnablePrintRoutes: true,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return mw.Error(ctx, err)
		},
	})

	app.Use(cors.New())
	app.Use(recover.New())

	return app
}
