package entity

import (
	"slices"
)

type Symbol struct {
	Model
	Code string `json:"code"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Symbols []Symbol

func (m *Symbols) Exist(symbol string) bool {
	return slices.ContainsFunc(*m, func(s Symbol) bool {
		return s.Code == symbol
	})
}
