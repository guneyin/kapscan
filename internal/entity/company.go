package entity

import (
	"gorm.io/gorm"
	"slices"
	"time"
)

type Company struct {
	gorm.Model
	Code     string `gorm:"uniqueIndex"`
	MemberID string
	Name     string
	Address  string
	Email    string
	Website  string
	Index    string
	Sector   string
	Market   string
	Icon     string
	Shares   CompanyShareList
}

func (c *Company) AddShare(cs CompanyShare) {
	c.Shares = append(c.Shares, cs)
}

type CompanyList []Company

func (cl *CompanyList) Exist(code string) bool {
	return slices.ContainsFunc(*cl, func(c Company) bool {
		return c.Code == code
	})
}

type CompanyShareList []CompanyShare

type CompanyShare struct {
	CompanyID       uint
	Date            time.Time `gorm:"index"`
	Title           string
	CapitalByAmount float64
	CapitalByVolume float64
	VoteRight       float64
}

func (c *Company) BeforeSave(tx *gorm.DB) (err error) {
	return tx.Where("company_id = ?", c.ID).Delete(&CompanyShare{}).Error
}
