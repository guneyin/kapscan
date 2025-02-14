package entity

import (
	"time"

	"gorm.io/gorm"
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

type CompanyShare struct {
	gorm.Model
	CompanyID    uint
	Date         time.Time            `gorm:"index"`
	ShareHolders []CompanyShareHolder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CompanyShareHolder struct {
	CompanyShareID  uint
	Title           string
	CapitalByAmount float64
	CapitalByVolume float64
	VoteRight       float64
}

type (
	CompanyList      []Company
	CompanyShareList []CompanyShare
)

func (c *Company) AddShareHolder(dt time.Time, csh CompanyShareHolder) {
	for i, share := range c.Shares {
		if share.Date == dt || share.Date.IsZero() {
			c.Shares[i].ShareHolders = append(c.Shares[i].ShareHolders, csh)
			return
		}
	}
	c.Shares = append(c.Shares, CompanyShare{
		Date:         dt,
		ShareHolders: make([]CompanyShareHolder, 0),
	})
}

func (c *Company) BeforeSave(tx *gorm.DB) error {
	return tx.Where("company_id = ?", c.ID).Delete(&CompanyShare{}).Error
}
