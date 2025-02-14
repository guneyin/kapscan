package company

import (
	"context"

	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/company"
	"github.com/guneyin/kapscan/internal/util"
	"github.com/vcraescu/go-paginator/v2"
)

type Pager struct {
	repo   *company.Repo
	search string
	offset int
	limit  int
	pg     paginator.Paginator
}

func (lg *Pager) Offset(offset int) *Pager {
	lg.offset = offset
	return lg
}

func (lg *Pager) Limit(limit int) *Pager {
	lg.limit = limit
	return lg
}

func (lg *Pager) Do(ctx context.Context) (*Pager, error) {
	pg, err := lg.repo.Search(ctx, lg.search, lg.offset, lg.limit)
	if err != nil {
		return nil, err
	}
	lg.pg = pg

	return lg, nil
}

func (lg *Pager) PageData() paginator.Paginator {
	return lg.pg
}

func (lg *Pager) Data() (*entity.CompanyList, error) {
	data := &entity.CompanyList{}
	err := lg.pg.Results(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (lg *Pager) DataAs(m any) error {
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
