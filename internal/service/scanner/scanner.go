package scanner

import (
	"context"

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

func (s *Service) GetCompanyList() (entity.CompanyList, error) {
	return s.repo.FetchCompanyList()
}

func (s *Service) SyncCompanyList(ctx context.Context) error {
	companySvc := company.NewService()

	companyList, err := s.GetCompanyList()
	if err != nil {
		return err
	}

	dbCompanyList, err := companySvc.GetAll()
	if err != nil {
		return err
	}

	for _, cmp := range companyList {
		if !dbCompanyList.Exist(cmp.Code) {
			fetched, err := s.repo.FetchCompany(ctx, cmp.Code)
			if err != nil {
				return err
			}

			err = companySvc.Save(fetched)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) SyncCompany(ctx context.Context, cmp *entity.Company) error {
	fetched, err := s.repo.FetchCompany(ctx, cmp.Code)
	if err != nil {
		return err
	}
	cmp = fetched

	return nil
}
