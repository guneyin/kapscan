package dto

import (
	"slices"
	"strings"
	"time"
)

type Company struct {
	Code    string         `json:"code"`
	Name    string         `json:"name"`
	Address string         `json:"address"`
	Email   string         `json:"email"`
	Website string         `json:"website"`
	Index   string         `json:"index"`
	Sector  string         `json:"sector"`
	Market  string         `json:"market"`
	Icon    string         `json:"icon"`
	Price   string         `json:"price"`
	Shares  []CompanyShare `json:"shares"`
}

type CompanyShare struct {
	Date         time.Time            `json:"date"`
	ShareHolders []CompanyShareHolder `json:"shareholders"`
}

type CompanyShareHolder struct {
	Title           string  `json:"title"`
	CapitalByAmount float64 `json:"capitalByAmount"`
	CapitalByVolume float64 `json:"capitalByVolume"`
	VoteRight       float64 `json:"voteRight"`
}

type CompanyList []Company

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

func (cl CompanyList) Exist(code string) bool {
	return slices.ContainsFunc(cl, func(c Company) bool {
		return c.Code == code
	})
}
