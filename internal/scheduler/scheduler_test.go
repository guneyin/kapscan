package scheduler

import (
	"github.com/guneyin/kapscan/internal/store"
	"os"
	"path"
	"runtime"
	"testing"
)

func changeWorkDir() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func Test_Cron(t *testing.T) {
	changeWorkDir()
	_ = store.InitDB(store.DBTest)

	syncCompanyList()
	syncCompanyInfo()
}
