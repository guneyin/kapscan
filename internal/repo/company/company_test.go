package company

import (
	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_GetCompanyList(t *testing.T) {
	_ = store.InitDB(store.DBTest)

	repo := NewRepo()

	companyList := dto.CompanyList{}
	data, err := repo.Search("", -1, -1)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	err = data.Results(companyList)
	assert.NoError(t, err)

	for _, c := range companyList {
		t.Log(c.Name)
	}
}
