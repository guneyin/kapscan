package scanner

import (
	"context"

	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/logger"
	"github.com/guneyin/kapscan/internal/repo/scanner"
	"github.com/guneyin/kapscan/internal/service/company"
)

type Service struct {
	repo *scanner.Repo
}

func NewService() *Service {
	return &Service{
		repo: scanner.NewRepo(),
	}
}

func (s *Service) GetCompanyList(ctx context.Context) (entity.CompanyList, error) {
	return s.repo.FetchCompanyList(ctx)
}

func (s *Service) SyncCompanyList(ctx context.Context) error {
	companySvc := company.NewService()

	companyList, err := s.GetCompanyList(ctx)
	if err != nil {
		return err
	}

	dbCompanyList, err := companySvc.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, cmp := range companyList {
		if !dbCompanyList.Exist(cmp.Code) {
			err = s.SyncCompanyWithShares(ctx, &cmp)
			if err != nil {
				logger.Log().ErrorContext(ctx, "sync company %s error: %v", cmp.Code, err)
				continue
			}
		}
	}

	return nil
}

func (s *Service) SyncCompanyWithShares(ctx context.Context, cmp *entity.Company) error {
	companySvc := company.NewService()

	err := s.repo.SyncCompanyWithShares(ctx, cmp)
	if err != nil {
		return err
	}

	return companySvc.Save(ctx, cmp)
}
