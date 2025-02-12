package util

import (
	"os"
	"path"
	"runtime"
)

func ChangeWorkDir() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
