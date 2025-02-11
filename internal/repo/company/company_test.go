package company_test

import (
	"testing"

	"github.com/guneyin/kapscan/internal/repo/company"

	"github.com/guneyin/kapscan/internal/dto"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepo_GetCompanyList(t *testing.T) {
	_ = store.InitDB(store.DBTest)

	repo := company.NewRepo()

	companyList := dto.CompanyList{}
	data, err := repo.Search("", -1, -1)
	require.NoError(t, err)
	assert.NotEmpty(t, data)

	err = data.Results(companyList)
	require.NoError(t, err)

	for _, c := range companyList {
		t.Log(c.Name)
	}
}
