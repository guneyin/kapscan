package scheduler

import (
	"context"
	"github.com/guneyin/kapscan/internal/service/company"
	"github.com/guneyin/kapscan/internal/service/scanner"
	"github.com/robfig/cron"
	"log"
)

type Cron struct {
	cron *cron.Cron
}

func New() (*Cron, func()) {
	c := cron.New()
	return &Cron{cron: c}, c.Stop
}

func (c *Cron) Start() {
	_ = c.AddJob("@every 00h30m00s", syncCompanyList)
	go c.cron.Start()
}

func (c *Cron) AddJob(spec string, cmd func()) error {
	return c.cron.AddFunc(spec, cmd)
}

func syncCompanyList() {
	ctx := context.Background()

	log.Printf("sync company list started")

	svc := scanner.NewService()
	err := svc.SyncCompanyList(ctx)
	if err != nil {
		log.Println(err)
	}
}

func syncCompanyInfo() {
	log.Printf("sync company info started")

	scannerSvc := scanner.NewService()
	companySvc := company.NewService()

	cl, err := companySvc.GetAll()
	if err != nil {
		return
	}

	for _, comp := range cl {
		_ = scannerSvc.SyncCompany(context.Background(), &comp)
		err = companySvc.Save(&comp)
		if err != nil {
			return
		}
	}
}
