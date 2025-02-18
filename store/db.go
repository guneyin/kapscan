package store

import (
	"database/sql"
	"errors"
	"os"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type DBEnvironment string

var (
	DBProd DBEnvironment = "prod"
	DBTest DBEnvironment = "test"
)

var (
	once sync.Once
	db   *bun.DB

	ErrUnableToConnectToDatabase = errors.New("unable to connect to database")
	ErrDatabaseMigrationFailed   = errors.New("database migration failed")
	ErrDatabaseNil               = errors.New("database is nil")
	ErrInvalidDatabaseProvider   = errors.New("invalid database provider")
)

const dataDir = "./data/data.db"

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

func Get() *bun.DB {
	if db == nil {
		panic(ErrDatabaseNil)
	}

	return db
}

func initDBFile() {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		file, _ := os.Create(dataDir)
		defer file.Close()
	}
}

func newProdDB() (*bun.DB, error) {
	initDBFile()

	sqlDB, err := sql.Open(sqliteshim.DriverName(), dataDir)
	if err != nil {
		panic(err)
	}

	bunDB := bun.NewDB(sqlDB, sqlitedialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	//err = migrate(sqlDB)
	//if err != nil {
	//	return nil, err
	//}

	return bunDB, nil
}

func newTestDB() (*bun.DB, error) {
	sqlDB, err := sql.Open(sqliteshim.DriverName(), "file::memory:?cache=shared")
	if err != nil {
		return nil, err
	}

	bunDB := bun.NewDB(sqlDB, sqlitedialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	//err = migrate(sqlDB)
	//if err != nil {
	//	return nil, err
	//}

	return bunDB, nil
}

//func migrate(sqlDB *sql.DB) error {
//	if err := goose.SetDialect("sqlite"); err != nil {
//		return err
//	}
//
//	if err := goose.Up(sqlDB, "migrations"); err != nil {
//		logger.Log().Error(err.Error())
//	}
//
//	return nil
//}
