package migration

import (
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

func Init() error {
	return Migrations.DiscoverCaller()
}
