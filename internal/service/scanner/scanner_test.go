package scanner_test

import (
	"context"
	"github.com/guneyin/kapscan/internal/entity"
	"github.com/guneyin/kapscan/internal/service/scanner"
	"github.com/guneyin/kapscan/internal/store"
	"github.com/guneyin/kapscan/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_SyncCompany(t *testing.T) {
	util.ChangeWorkDir()
	_ = store.InitDB(store.DBProd)

	svc := scanner.NewService()

	err := svc.SyncCompany(context.Background(), &entity.Company{Code: "TUPRS"})
	assert.Nil(t, err)
}
