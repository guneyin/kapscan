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

func (c *Company) AddShareHolder(dt time.Time, csh CompanyShareHolder) {
	for i, share := range c.Shares {
		if share.Date == dt || share.Date.IsZero() {
			c.Shares[i].ShareHolders = append(c.Shares[i].ShareHolders, csh)
			//share.ShareHolders = append(share.ShareHolders, csh)
			return
		}
	}
	c.Shares = append(c.Shares, CompanyShare{
		Date:         dt,
		ShareHolders: make(CompanyShareHolderList, 0),
	})
}

type CompanyList []Company

func (cl *CompanyList) Exist(code string) bool {
	return slices.ContainsFunc(*cl, func(c Company) bool {
		return c.Code == code
	})
}

type CompanyShareList []CompanyShare

type CompanyShare struct {
	gorm.Model
	CompanyID    uint
	Date         time.Time              `gorm:"index"`
	ShareHolders CompanyShareHolderList `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CompanyShareHolderList []CompanyShareHolder

type CompanyShareHolder struct {
	CompanyShareID  uint
	Title           string
	CapitalByAmount float64
	CapitalByVolume float64
	VoteRight       float64
}

func (c *Company) BeforeSave(tx *gorm.DB) (err error) {
	return tx.Where("company_id = ?", c.ID).Delete(&CompanyShare{}).Error
}
