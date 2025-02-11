package company

import (
	"github.com/guneyin/gobist"
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
	pg     paginator.Paginator
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

	price := ""
	bist := gobist.New()
	q, _ := bist.GetQuote([]string{c.Code})
	if q != nil {
		price = q.Items[0].Price
	}

	return util.Convert(c, &dto.Company{Price: price})
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

func (lg *ListGetter) Do() (*ListGetter, error) {
	pg, err := lg.repo.GetCompanyList(lg.search, lg.offset, lg.limit)
	if err != nil {
		return nil, err
	}
	lg.pg = pg

	return lg, nil
}

func (lg *ListGetter) PageData() paginator.Paginator {
	return lg.pg
}

func (lg *ListGetter) Data() (*entity.CompanyList, error) {
	data := &entity.CompanyList{}
	err := lg.pg.Results(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (lg *ListGetter) DataAs(m any) error {
	data, err := lg.Data()
	if err != nil {
		return err
	}

	_, err = util.Convert(data, m)
	if err != nil {
		return err
	}
	return nil
}
