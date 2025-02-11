package entity

import (
	"slices"
	"time"

	"github.com/oklog/ulid/v2"
)

type Company struct {
	Model
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
	Model
	CompanyID       ulid.ULID `json:"companyID"`
	Date            time.Time `json:"date" gorm:"index"`
	Title           string    `json:"title"`
	CapitalByAmount float64   `json:"capitalByAmount"`
	CapitalByVolume float64   `json:"capitalByVolume"`
	VoteRight       float64   `json:"voteRight"`
}
