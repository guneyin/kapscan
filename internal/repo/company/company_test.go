package company

import (
	"github.com/guneyin/kapscan/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_GetCompanyList(t *testing.T) {
	_ = store.InitDB(store.DBTest)

	repo := NewRepo()
	companyList, _, err := repo.GetCompanyList("", -1, -1)
	assert.NoError(t, err)
	assert.NotEmpty(t, companyList)

	for _, c := range companyList {
		t.Log(c.Name)
	}
}
