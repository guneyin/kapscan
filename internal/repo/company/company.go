package company

import (
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/scraper"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/oklog/ulid/v2"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
	"gorm.io/gorm/clause"
)

type Repo struct {
	scraper *scraper.Scraper
}

func NewRepo() *Repo {
	return &Repo{scraper: scraper.New()}
}

func (r *Repo) GetCompanyList(page, size int16) (entity.CompanyList, paginator.Paginator, error) {
	db := store.Get()

	var companyList entity.CompanyList
	stmt := db.Model(&entity.Company{})

	p := paginator.New(adapter.NewGORMAdapter(stmt), int(size))

	p.SetPage(int(page))
	err := p.Results(&companyList)
	if err != nil {
		return nil, nil, err
	}

	return companyList, p, nil
}

func (r *Repo) SaveCompany(company *entity.Company) error {
	db := store.Get()

	tx := db.Clauses(clause.OnConflict{UpdateAll: true}).Save(company)
	return tx.Error
}

func (r *Repo) GetCompany(id string) (*entity.Company, error) {
	uid, err := ulid.Parse(id)
	if err != nil {
		return nil, err
	}

	db := store.Get()

	company := &entity.Company{Model: entity.Model{ID: uid}}
	tx := db.First(&company)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return company, nil
}
