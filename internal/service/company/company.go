package company

import (
	"github.com/guneyin/gobist"
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/company"
	"github.com/guneyin/kapscan/util"
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

func (s *Service) GetAll() (entity.CompanyList, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByCode(code string) (*dto.Company, error) {
	cmp, err := s.repo.GetByCode(code)
	if err != nil {
		return nil, err
	}

	price := ""
	bist := gobist.New()
	q, _ := bist.GetQuote([]string{code})
	if q != nil {
		price = q.Items[0].Price
	}

	return util.Convert(cmp, &dto.Company{Price: price})
}

func (s *Service) Save(company *entity.Company) error {
	return s.repo.Save(company)
}
