package company

import (
	"context"
	"github.com/uptrace/bun"
	"strings"

	"github.com/guneyin/gobist"

	"github.com/guneyin/kapscan/entity"
	"github.com/guneyin/kapscan/store"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
)

type Repo struct {
	bist *gobist.Bist
	db   *bun.DB
}

func NewRepo() *Repo {
	return &Repo{
		bist: gobist.New(),
		db:   store.Get(),
	}
}

func (r *Repo) Search(ctx context.Context, search string, page, size int) (paginator.Paginator, error) {
	db := store.Get()

	search = "%" + strings.ToUpper(search) + "%"
	data := &entity.CompanyList{}
	err := db.NewSelect().
		Model(data).
		Where("code like ? or name like ?", search, search).
		Offset((page - 1) * size).
		Limit(size).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	//stmt := db.Model(&entity.Company{}).Where("code like ? or name like ?", search, search)

	p := paginator.New(adapter.NewSliceAdapter(data), size)
	p.SetPage(page)

	return p, nil
}

func (r *Repo) GetAll(ctx context.Context) (entity.CompanyList, error) {
	db := store.Get()

	data := entity.CompanyList{}
	return data, db.NewSelect().
		Model(&data).
		Scan(ctx)
}

func (r *Repo) SaveCompany(ctx context.Context, company *entity.Company) error {
	//db := store.Get(ctx)
	//
	//tx := db.Clauses(clause.OnConflict{
	//	Columns: []clause.Column{
	//		{Table: "companies", Name: "code"},
	//		{Table: "company_shares", Name: "company_id"},
	//		{Table: "company_shares", Name: "date"},
	//	},
	//	UpdateAll: true}).
	//	Save(company)
	//
	//return tx.Error
	return nil
}

func (r *Repo) SaveCompanyShare(ctx context.Context, share *entity.CompanyShare) error {

}

func (r *Repo) GetByCode(ctx context.Context, code string) (*entity.Company, error) {
	db := store.Get()

	data := &entity.Company{}
	return data, db.NewSelect().
		Model(data).
		Where("code = ?", code).
		Scan(ctx)
}
