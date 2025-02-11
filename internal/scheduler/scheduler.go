package scheduler

import (
	"context"
	"github.com/guneyin/kapscan/internal/entity"
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
	log.Printf("sync company list started")

	scannerSvc := scanner.NewService()
	companySvc := company.NewService()

	companyList, err := scannerSvc.GetCompanyList()
	if err != nil {
		return
	}

	cl, err := companySvc.GetCompanyList().Do()
	if err != nil {
		return
	}

	dbCompanyList := entity.CompanyList{}
	err = cl.DataAs(dbCompanyList)
	if err != nil {
		return
	}

	for _, comp := range companyList {
		if !dbCompanyList.Exist(comp.Code) {
			_ = scannerSvc.SyncCompany(context.Background(), &comp)

			err = companySvc.SaveCompany(&comp)
			if err != nil {
				return
			}
		}
	}
}

func syncCompanyInfo() {
	log.Printf("sync company info started")

	scannerSvc := scanner.NewService()
	companySvc := company.NewService()

	cl, err := companySvc.GetCompanyList().Do()
	if err != nil {
		return
	}

	companyList, err := cl.Data()
	if err != nil {
		return
	}

	for _, comp := range *companyList {
		_ = scannerSvc.SyncCompany(context.Background(), &comp)
		err = companySvc.SaveCompany(&comp)
		if err != nil {
			return
		}
	}
}
