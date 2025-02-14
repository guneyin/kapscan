package scanner_test

import (
	"context"
	"testing"

	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/service/scanner"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/guneyin/kapscan/internal/util"
	"github.com/stretchr/testify/require"
)

func TestService_SyncCompany(t *testing.T) {
	util.ChangeWorkDir()
	_ = store.InitDB(store.DBProd)

	svc := scanner.NewService()

	err := svc.SyncCompanyWithShares(context.Background(), &entity.Company{Code: "TUPRS"})
	require.NoError(t, err)
}
