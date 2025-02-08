package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/guneyin/kapscan/internal/config"
	"github.com/guneyin/kapscan/internal/controller"
	"github.com/guneyin/kapscan/internal/scheduler"
	"github.com/guneyin/kapscan/internal/server"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/guneyin/kapscan/util"
	"log/slog"
	"os"
	"time"
)

const appName = "KAPScan"

type Application struct {
	Name       string
	Version    string
	Config     *config.Config
	Server     *fiber.App
	Controller *controller.Controller
}

func NewApplication(name string, cfg *config.Config, srv *fiber.App, cnt *controller.Controller) *Application {
	return &Application{
		Name:       name,
		Config:     cfg,
		Server:     srv,
		Controller: cnt,
	}
}

func main() {
	util.SetLastRun(time.Now())
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	cfg, err := config.NewConfig()
	checkError(err)

	err = store.InitDB(store.DBProd)
	checkError(err)

	srv := server.NewServer(appName)
	checkError(err)

	api := srv.Group("/api")
	cnt := controller.NewController(api)

	app := NewApplication(appName, cfg, srv, cnt)

	cron, stop := scheduler.New()
	defer stop()
	cron.Start()

	log.Fatal(app.Server.Listen(fmt.Sprintf(":%d", app.Config.HttpPort)))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
