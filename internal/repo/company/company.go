package company

import (
	"strings"

	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/scraper"
	"github.com/guneyin/kapscan/internal/store"
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

func (r *Repo) Search(search string, page, size int) (paginator.Paginator, error) {
	db := store.Get()

	search = "%" + strings.ToUpper(search) + "%"
	stmt := db.Model(&entity.Company{}).Where("code like ? or name like ?", search, search)

	p := paginator.New(adapter.NewGORMAdapter(stmt), size)
	p.SetPage(page)

	return p, nil
}

func (r *Repo) GetAll() (entity.CompanyList, error) {
	db := store.Get()

	var companies entity.CompanyList
	tx := db.Model(&entity.Company{}).Find(&companies)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return companies, nil
}

func (r *Repo) Save(company *entity.Company) error {
	db := store.Get()

	tx := db.Clauses(clause.OnConflict{UpdateAll: true}).Save(company)
	return tx.Error
}

func (r *Repo) GetByCode(code string) (*entity.Company, error) {
	db := store.Get()

	company := &entity.Company{}
	tx := db.Where("code = ?", code).First(&company)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return company, nil
}
