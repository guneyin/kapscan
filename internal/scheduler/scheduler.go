package scheduler

import (
	"github.com/guneyin/kapscan/internal/repo"
	"github.com/guneyin/kapscan/internal/service/scanner"
	"github.com/robfig/cron"
	"log"
)

type Cron struct {
	cron *cron.Cron
	repo *repo.Repo
}

func New() (*Cron, func()) {
	c := cron.New()
	return &Cron{cron: c}, c.Stop
}

func (c *Cron) Start() {
	_ = c.addJob("@every 00h00m15s", syncSymbolList)
	go c.cron.Start()
}

func (c *Cron) addJob(spec string, cmd func()) error {
	return c.cron.AddFunc(spec, cmd)
}

func syncSymbolList() {
	log.Printf("sync symbol list started")

	svc := scanner.NewScannerService()

	fetched, err := svc.FetchSymbolList()
	if err != nil {
		return
	}

	list, err := svc.GetSymbolList(-1, -1)
	if err != nil {
		return
	}

	for _, symbol := range fetched {
		if !list.Exist(symbol.Code) {
			err = svc.SaveSymbol(&symbol)
			if err != nil {
				return
			}
		}
	}
}
