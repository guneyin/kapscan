package company

import (
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/company"
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
	offset int
	limit  int
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

func (lg *ListGetter) Offset(offset int) *ListGetter {
	lg.offset = offset
	return lg
}

func (lg *ListGetter) Limit(limit int) *ListGetter {
	lg.limit = limit
	return lg
}

func (lg *ListGetter) Do() (entity.CompanyList, error) {
	return lg.repo.GetCompanyList(lg.offset, lg.limit)
}
