package entity

import (
	"slices"
)

type Company struct {
	Model
	Code string `json:"code"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type CompanyList []Company

func (cl *CompanyList) Exist(code string) bool {
	return slices.ContainsFunc(*cl, func(c Company) bool {
		return c.Code == code
	})
}
