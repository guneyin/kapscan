package scheduler

import (
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
	_ = c.addJob("@every 00h30m00s", syncCompanyList)
	go c.cron.Start()
}

func (c *Cron) addJob(spec string, cmd func()) error {
	return c.cron.AddFunc(spec, cmd)
}

func syncCompanyList() {
	log.Printf("sync symbol list started")

	scannerSvc := scanner.NewService()
	companySvc := company.NewService()

	fetched, err := scannerSvc.GetCompanyList()
	if err != nil {
		return
	}

	list, _, err := companySvc.GetCompanyList().Do()
	if err != nil {
		return
	}

	for _, symbol := range fetched {
		if !list.Exist(symbol.Code) {
			err = companySvc.SaveCompany(&symbol)
			if err != nil {
				return
			}
		}
	}
}
