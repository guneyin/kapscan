package company_test

import (
	"context"
	"testing"

	"github.com/guneyin/kapscan/repo/company"

	"github.com/guneyin/kapscan/dto"
	"github.com/guneyin/kapscan/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepo_GetCompanyList(t *testing.T) {
	ctx := context.Background()
	_ = store.InitDB(store.DBTest)

	repo := company.NewRepo()

	companyList := dto.CompanyList{}
	data, err := repo.Search(ctx, "", -1, -1)
	require.NoError(t, err)
	assert.NotEmpty(t, data)

	err = data.Results(companyList)
	require.NoError(t, err)

	for _, c := range companyList {
		t.Log(c.Name)
	}
}
