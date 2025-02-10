package entity

import (
	"github.com/oklog/ulid/v2"
	"slices"
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
	Shares   CompanyShare
}

type CompanyList []Company

func (cl *CompanyList) Exist(code string) bool {
	return slices.ContainsFunc(*cl, func(c Company) bool {
		return c.Code == code
	})
}

type CompanyShare struct {
	Model
	CompanyID       ulid.ULID `json:"companyID"`
	Title           string    `json:"title"`
	CapitalByAmount float64   `json:"capitalByAmount"`
	CapitalByVolume float64   `json:"capitalByVolume"`
	VoteRight       float64   `json:"voteRight"`
}
