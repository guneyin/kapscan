package scanner

import (
	"context"
	"github.com/guneyin/kapscan/internal/dto"
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

func (s *Service) GetCompany(ctx context.Context, symbol string) ([]dto.ShareHolder, error) {
	return s.repo.GetCompany(ctx, symbol)
}
