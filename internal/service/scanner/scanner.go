package scanner

import (
	"context"
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo"
)

type Service struct {
	repo *repo.Repo
}

func NewScannerService() *Service {
	return &Service{
		repo: repo.New(),
	}
}

func (s *Service) FetchSymbolList() (entity.Symbols, error) {
	return s.repo.FetchSymbolList()
}

func (s *Service) GetSymbolList(offset, limit int) (entity.Symbols, error) {
	return s.repo.GetSymbolList(offset, limit)
}

func (s *Service) Scan(ctx context.Context, symbol string) ([]dto.ShareHolder, error) {
	return s.repo.ScanSymbol(ctx, symbol)
}

func (s *Service) SaveSymbol(symbol *entity.Symbol) error {
	return s.repo.SaveSymbol(symbol)
}
