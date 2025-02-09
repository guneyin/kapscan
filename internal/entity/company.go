package entity

import (
	"github.com/oklog/ulid/v2"
	"slices"
	"strings"
)

type Company struct {
	Model
	Code     string       `gorm:"uniqueIndex" json:"code"`
	MemberID string       `json:"memberID"`
	Name     string       `json:"name"`
	Address  string       `json:"address"`
	Email    string       `json:"email"`
	Website  string       `json:"website"`
	Index    string       `json:"index"`
	Sector   string       `json:"sector"`
	Market   string       `json:"market"`
	Icon     string       `json:"icon"`
	Shares   CompanyShare `json:"shares"`
}

func (c Company) AvatarText() string {
	parts := strings.Split(c.Name, " ")

	var at strings.Builder
	for i, part := range parts {
		if i > 1 {
			break
		}

		letter := strings.ToUpper(string(part[0]))
		at.WriteString(letter)
	}

	return at.String()
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
