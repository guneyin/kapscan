package scheduler

import (
	"context"
	"time"

	"github.com/guneyin/kapscan/internal/logger"

	"github.com/guneyin/kapscan/internal/service/company"
	"github.com/guneyin/kapscan/internal/service/scanner"
	"github.com/robfig/cron"
)

type Cron struct {
	cron *cron.Cron
}

func New() (*Cron, func()) {
	c := cron.New()
	return &Cron{cron: c}, c.Stop
}

func (c *Cron) Start() {
	_ = c.AddJob("@every 24h00m00s", SyncCompanyList)
	go c.cron.Start()
}

func (c *Cron) AddJob(spec string, cmd func()) error {
	return c.cron.AddFunc(spec, cmd)
}

func SyncCompanyList() {
	ctx, closer := context.WithTimeout(context.Background(), 10*time.Minute)
	defer closer()

	logger.Log().InfoContext(ctx, "sync company list started")

	scannerSvc := scanner.NewService()
	err := scannerSvc.SyncCompanyList(ctx)
	if err != nil {
		logger.Log().ErrorContext(ctx, err.Error())
		return
	}
}

func SyncCompanyInfo() {
	ctx, closer := context.WithTimeout(context.Background(), 10*time.Minute)
	defer closer()

	logger.Log().InfoContext(ctx, "sync company info started")

	scannerSvc := scanner.NewService()
	companySvc := company.NewService()

	cl, err := companySvc.GetAll(ctx)
	if err != nil {
		return
	}

	for _, cmp := range cl {
		err = scannerSvc.SyncCompanyWithShares(context.Background(), &cmp)
		if err != nil {
			logger.Log().ErrorContext(ctx, err.Error())
			return
		}

		err = companySvc.Save(ctx, &cmp)
		if err != nil {
			logger.Log().ErrorContext(ctx, err.Error())
			return
		}
	}
}
