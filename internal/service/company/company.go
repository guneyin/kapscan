package company

import (
	"context"

	"github.com/guneyin/gobist"
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/util"

	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/company"
)

type Service struct {
	repo *company.Repo
}

func NewService() *Service {
	return &Service{
		repo: company.NewRepo(),
	}
}

func (s *Service) Search(term string) *Pager {
	return &Pager{
		repo:   s.repo,
		search: term,
		offset: -1,
		limit:  -1,
	}
}

func (s *Service) GetAll(ctx context.Context) (dto.CompanyList, error) {
	companyList, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return util.Convert(companyList, dto.CompanyList{})
}

func (s *Service) GetByCode(ctx context.Context, code string) (*dto.Company, error) {
	cmp, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	price := ""
	bist := gobist.New()
	q, _ := bist.GetQuote(ctx, code)
	if q != nil {
		price = q.Price
	}

	return util.Convert(cmp, &dto.Company{Price: price})
}

func (s *Service) Save(ctx context.Context, company *entity.Company) error {
	return s.repo.Save(ctx, company)
}
