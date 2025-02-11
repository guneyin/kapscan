package server

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/guneyin/kapscan/internal/mw"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func NewServer(appName string) *fiber.App {
	engine := html.New("./web/views", ".html")

	app := fiber.New(fiber.Config{
		ServerHeader:      fmt.Sprintf("%s HTTP Server", appName),
		BodyLimit:         16 * 1024 * 1024,
		AppName:           appName,
		EnablePrintRoutes: true,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		ErrorHandler:      errorHandler,
		Views:             engine,
		//ViewsLayout:       "layouts/main",
	})

	app.Use(cors.New())
	app.Use(recover.New())

	return app
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	return mw.Error(ctx, err)
}
