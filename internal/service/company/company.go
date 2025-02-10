package company

import (
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/company"
	"github.com/guneyin/kapscan/util"
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
	search string
	offset int16
	limit  int16
}

func (s *Service) GetCompanyList() *ListGetter {
	return &ListGetter{
		repo:   s.repo,
		search: "",
		offset: -1,
		limit:  -1,
	}
}

func (s *Service) GetCompany(id string) (*dto.Company, error) {
	c, err := s.repo.GetCompany(id)
	if err != nil {
		return nil, err
	}

	return util.Convert(c, &dto.Company{})
}

func (s *Service) SaveCompany(company *entity.Company) error {
	return s.repo.SaveCompany(company)
}

func (lg *ListGetter) Search(s string) *ListGetter {
	lg.search = s
	return lg
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
	return lg.repo.GetCompanyList(lg.search, lg.offset, lg.limit)
}
