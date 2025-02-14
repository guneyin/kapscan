package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/internal/config"
	"github.com/guneyin/kapscan/internal/controller"
	"github.com/guneyin/kapscan/internal/scheduler"
	"github.com/guneyin/kapscan/internal/server"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/guneyin/kapscan/internal/util"
	"github.com/guneyin/kapscan/web/handler"
)

const appName = "KAPScan"

type Application struct {
	Name       string
	Version    string
	Config     *config.Config
	Server     *fiber.App
	Controller *controller.Controller
	WebHandler *handler.Handler
}

func NewApplication(
	name string,
	cfg *config.Config,
	srv *fiber.App,
	cnt *controller.Controller,
	webHnd *handler.Handler) *Application {
	return &Application{
		Name:       name,
		Config:     cfg,
		Server:     srv,
		Controller: cnt,
		WebHandler: webHnd,
	}
}

func main() {
	util.SetLastRun(time.Now())

	cfg, err := config.NewConfig()
	checkError(err)

	err = store.InitDB(store.DBProd)
	checkError(err)

	srv := server.NewServer(appName)
	checkError(err)

	api := srv.Group("/api")
	apiCnt := controller.NewController(api)

	webHnd := handler.NewWebHandler(srv)

	app := NewApplication(appName, cfg, srv, apiCnt, webHnd)

	cron, stop := scheduler.New()
	defer stop()
	cron.Start()

	err = app.Server.Listen(fmt.Sprintf(":%d", app.Config.HTTPPort))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
