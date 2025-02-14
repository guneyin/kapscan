package scheduler_test

import (
	"testing"

	"github.com/guneyin/kapscan/internal/util"

	"github.com/guneyin/kapscan/internal/scheduler"

	"github.com/guneyin/kapscan/internal/store"
)

func Test_Cron(_ *testing.T) {
	util.ChangeWorkDir()
	_ = store.InitDB(store.DBTest)

	scheduler.SyncSymbolList()
	scheduler.SyncCompany()
}
