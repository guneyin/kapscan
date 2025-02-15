package company

import (
	"context"
	"strings"

	"github.com/guneyin/gobist"

	"gorm.io/gorm/clause"

	"github.com/guneyin/kapscan/entity"
	"github.com/guneyin/kapscan/store"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
)

type Repo struct {
	bist *gobist.Bist
}

func NewRepo() *Repo {
	return &Repo{bist: gobist.New()}
}

func (r *Repo) Search(ctx context.Context, search string, page, size int) (paginator.Paginator, error) {
	db := store.Get(ctx)

	search = "%" + strings.ToUpper(search) + "%"
	stmt := db.Model(&entity.Company{}).Where("code like ? or name like ?", search, search)

	p := paginator.New(adapter.NewGORMAdapter(stmt), size)
	p.SetPage(page)

	return p, nil
}

func (r *Repo) GetAll(ctx context.Context) (entity.CompanyList, error) {
	db := store.Get(ctx)

	var companies entity.CompanyList
	tx := db.Model(&entity.Company{}).Find(&companies)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return companies, nil
}

func (r *Repo) Save(ctx context.Context, company *entity.Company) error {
	db := store.Get(ctx)

	tx := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Table: "companies", Name: "code"},
			{Table: "company_shares", Name: "company_id"},
			{Table: "company_shares", Name: "date"},
		},
		UpdateAll: true}).
		Save(company)

	return tx.Error
}

func (r *Repo) GetByCode(ctx context.Context, code string) (*entity.Company, error) {
	db := store.Get(ctx)

	company := &entity.Company{}
	tx := db.Where("code = ?", code).
		Preload("Shares").
		Preload("Shares.ShareHolders").
		First(&company)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return company, nil
}
