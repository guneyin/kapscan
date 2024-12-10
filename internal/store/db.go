package store

import (
	"github.com/guneyin/kapscan/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path"
)

var gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Error)}

func NewDB(_ *config.Config) (*gorm.DB, error) {
	const dataDir = "data/data.db"

	if _, err := os.Stat(dataDir); err != nil {

		err = os.MkdirAll(path.Dir(dataDir), 0777)
		if err != nil {
			return nil, err
		}
	}

	return gorm.Open(sqlite.Open(dataDir), gormConfig)
}

func NewTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}
