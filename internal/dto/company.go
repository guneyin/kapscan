package dto

import (
	"strings"
	"time"
)

type Company struct {
	Code    string           `json:"code"`
	Name    string           `json:"name"`
	Address string           `json:"address"`
	Email   string           `json:"email"`
	Website string           `json:"website"`
	Index   string           `json:"index"`
	Sector  string           `json:"sector"`
	Market  string           `json:"market"`
	Icon    string           `json:"icon"`
	Price   string           `json:"price"`
	Shares  CompanyShareList `json:"shares"`
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

type CompanyShareList []CompanyShare

type CompanyShare struct {
	Date            time.Time `json:"date"`
	Title           string    `json:"title"`
	CapitalByAmount float64   `json:"capitalByAmount"`
	CapitalByVolume float64   `json:"capitalByVolume"`
	VoteRight       float64   `json:"voteRight"`
}
