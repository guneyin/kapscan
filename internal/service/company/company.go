package company

import (
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/company"
	"github.com/vcraescu/go-paginator/v2"
)

type Service struct {
	repo *company.Repo
	lg   *ListGetter
}

func NewService() *Service {
	return &Service{
		repo: company.NewRepo(),
	}
}

type ListGetter struct {
	repo   *company.Repo
	offset int16
	limit  int16
}

func (s *Service) GetCompanyList() *ListGetter {
	return &ListGetter{
		repo:   s.repo,
		offset: -1,
		limit:  -1,
	}
}

func (s *Service) GetCompany(id string) (*entity.Company, error) {
	return s.repo.GetCompany(id)
}

func (s *Service) SaveCompany(company *entity.Company) error {
	return s.repo.SaveCompany(company)
}

func (lg *ListGetter) Offset(offset int16) *ListGetter {
	lg.offset = offset
	return lg
}

func (lg *ListGetter) Limit(limit int16) *ListGetter {
	lg.limit = limit
	return lg
}

func (lg *ListGetter) Do() (entity.CompanyList, paginator.Paginator, error) {
	return lg.repo.GetCompanyList(lg.offset, lg.limit)
}
