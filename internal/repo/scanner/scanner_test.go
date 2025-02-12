package scanner_test

import (
	"context"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/repo/scanner"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepo_FetchCompany(t *testing.T) {
	ctx := context.Background()
	repo := scanner.NewRepo()

	cmp, err := repo.SyncCompany(ctx, &entity.Company{Code: "AFYON"})
	require.NoError(t, err)
	require.NotNil(t, cmp)
}
