package dto

import (
	"github.com/oklog/ulid/v2"
	"strings"
)

type Company struct {
	ID      ulid.ULID `json:"id"`
	Code    string    `json:"code"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Email   string    `json:"email"`
	Website string    `json:"website"`
	Index   string    `json:"index"`
	Sector  string    `json:"sector"`
	Market  string    `json:"market"`
	Icon    string    `json:"icon"`
	Price   string    `json:"price"`
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
