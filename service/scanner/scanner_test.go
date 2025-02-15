package scanner_test

import (
	"context"
	"testing"

	"github.com/guneyin/kapscan/entity"
	"github.com/guneyin/kapscan/service/scanner"
	"github.com/guneyin/kapscan/store"
	"github.com/guneyin/kapscan/util"
	"github.com/stretchr/testify/require"
)

func TestService_SyncCompany(t *testing.T) {
	util.ChangeWorkDir()
	_ = store.InitDB(store.DBProd)

	svc := scanner.NewService()

	err := svc.SyncCompanyWithShares(context.Background(), &entity.Company{Code: "TUPRS"})
	require.NoError(t, err)
}
