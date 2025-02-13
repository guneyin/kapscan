package scanner_test

import (
	"context"
	"testing"

	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/scanner"
	"github.com/stretchr/testify/require"
)

func TestRepo_FetchCompany(t *testing.T) {
	ctx := context.Background()
	repo := scanner.NewRepo()

	cmp := &entity.Company{Code: "AFYON"}
	err := repo.SyncCompany(ctx, cmp)
	require.NoError(t, err)
	require.NotNil(t, cmp)
}
