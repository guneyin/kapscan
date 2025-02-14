package scanner

import (
	"context"

	"github.com/guneyin/kapscan/internal/logger"

	"github.com/guneyin/kapscan/internal/entity"
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

func (s *Service) SyncCompanyWithShares(ctx context.Context, cmp *entity.Company) error {
	companySvc := company.NewService()

	err := s.repo.SyncCompanyWithShares(ctx, cmp)
	if err != nil {
		return err
	}

	return companySvc.Save(ctx, cmp)
}

func (s *Service) SyncSymbolList(ctx context.Context, limit int) error {
	companySvc := company.NewService()

	companyList, err := companySvc.GetAll(ctx)
	if err != nil {
		return err
	}

	symbolList, err := s.repo.GetSymbolList(ctx)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = len(companyList)
	}

	cnt := 0
	for _, cmp := range symbolList {
		if !companyList.Exist(cmp.Code) {
			if cnt > limit {
				break
			}
			cnt++

			err = companySvc.Save(ctx, &cmp)
			if err != nil {
				logger.Log().ErrorContext(ctx, err.Error())
			}
		}
	}

	return nil
}
