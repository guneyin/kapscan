package scanner

import (
	"context"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/scanner"
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
	return s.repo.GetCompanyList()
}

func (s *Service) SyncCompany(ctx context.Context, cmp *entity.Company) error {
	return s.repo.SyncCompany(ctx, cmp)
}
