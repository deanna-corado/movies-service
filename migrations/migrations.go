package migrations

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
)

func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		CreateMoviesTableMigration(),
		CreateYearColumnMigration(),
	}
}


