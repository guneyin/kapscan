package scanner_test

import (
	"context"
	"github.com/guneyin/kapscan/internal/repo/scanner"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepo_FetchCompany(t *testing.T) {
	ctx := context.Background()
	repo := scanner.NewRepo()

	cmp, err := repo.FetchCompany(ctx, "AFYON")
	require.NoError(t, err)
	require.NotNil(t, cmp)
}
