package repo

import (
	"github.com/guneyin/kapscan/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_GetSymbolList(t *testing.T) {
	_ = store.InitDB(store.DBTest)

	repo := New()
	symbolList, err := repo.GetSymbolList(-1, -1)
	assert.NoError(t, err)
	assert.NotEmpty(t, symbolList)

	for _, symbol := range symbolList {
		t.Log(symbol.Name)
	}
}
