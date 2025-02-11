package scheduler_test

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/guneyin/kapscan/internal/scheduler"

	"github.com/guneyin/kapscan/internal/store"
)

func changeWorkDir() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func Test_Cron(_ *testing.T) {
	changeWorkDir()
	_ = store.InitDB(store.DBTest)

	scheduler.SyncCompanyList()
	scheduler.SyncCompanyInfo()
}
