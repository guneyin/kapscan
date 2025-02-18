package main

import (
	"fmt"
	"github.com/guneyin/kapscan/migration"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/guneyin/kapscan/config"
	"github.com/guneyin/kapscan/controller"
	"github.com/guneyin/kapscan/scheduler"
	"github.com/guneyin/kapscan/server"
	"github.com/guneyin/kapscan/store"
	"github.com/guneyin/kapscan/util"
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

	err := store.InitDB(store.DBProd)
	checkError(err)

	err = migration.Init()
	checkError(err)

	app := &cli.App{
		Name: "kapscan",
		Commands: []*cli.Command{
			serveCmd(),
			migrateCMD(migrate.NewMigrator(store.Get(), migration.Migrations)),
		},
	}

	if err = app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func serveCmd() *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "start a kapscan server",
		Action: func(c *cli.Context) error {
			cfg, err := config.NewConfig()
			if err != nil {
				return err
			}

			srv := server.NewServer(appName)

			api := srv.Group("/api")
			apiCnt := controller.NewController(api)

			webHnd := handler.NewWebHandler(srv)

			app := NewApplication(appName, cfg, srv, apiCnt, webHnd)

			cron, stop := scheduler.New()
			defer stop()
			cron.Start()

			return app.Server.Listen(fmt.Sprintf(":%d", app.Config.HTTPPort))
		},
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
