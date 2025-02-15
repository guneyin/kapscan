package scheduler_test

import (
	"testing"

	"github.com/guneyin/kapscan/util"

	"github.com/guneyin/kapscan/scheduler"

	"github.com/guneyin/kapscan/store"
)

func Test_Cron(_ *testing.T) {
	util.ChangeWorkDir()
	_ = store.InitDB(store.DBTest)

	scheduler.SyncSymbolList()
	scheduler.SyncCompany()
}
