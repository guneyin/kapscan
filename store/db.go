package store

import (
	"context"
	"errors"
	"os"
	"path"
	"sync"

	"github.com/guneyin/kapscan/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBEnvironment string

var (
	DBProd DBEnvironment = "prod"
	DBTest DBEnvironment = "test"
)

var (
	once sync.Once
	db   *gorm.DB

	ErrUnableToConnectToDatabase = errors.New("unable to connect to database")
	ErrDatabaseMigrationFailed   = errors.New("database migration failed")
	ErrDatabaseNil               = errors.New("database is nil")
	ErrInvalidDatabaseProvider   = errors.New("invalid database provider")
)

func InitDB(env DBEnvironment) error {
	var err error
	once.Do(func() {
		switch env {
		case DBProd:
			db, err = newProdDB()
		case DBTest:
			db, err = newTestDB()
		default:
			err = ErrInvalidDatabaseProvider
		}
	})

	return err
}

func Get(ctx context.Context) *gorm.DB {
	if db == nil {
		panic(ErrDatabaseNil)
	}

	return db.WithContext(ctx)
}

var gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Error)}

func newProdDB() (*gorm.DB, error) {
	const dataDir = "data/data.db"
	if _, err := os.Stat(dataDir); err != nil {
		err = os.MkdirAll(path.Dir(dataDir), 0777)
		if err != nil {
			return nil, err
		}
	}

	gdb, err := gorm.Open(sqlite.Open(dataDir), gormConfig)
	if err != nil {
		return nil, err
	}

	return migrate(gdb)
}

func newTestDB() (*gorm.DB, error) {
	gdb, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
	if err != nil {
		return nil, errors.Join(ErrUnableToConnectToDatabase, err)
	}

	return migrate(gdb)
}

func migrate(gdb *gorm.DB) (*gorm.DB, error) {
	err := gdb.AutoMigrate(&entity.Company{}, &entity.CompanyShare{}, &entity.CompanyShareHolder{})
	if err != nil {
		return nil, errors.Join(ErrDatabaseMigrationFailed, err)
	}
	return gdb, nil
}
