package migrations

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateYearColumnMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "00002_create_year_column",
		Migrate: func(tx *gorm.DB) error {
			type Movie struct {
				Year int
			}
			return tx.Migrator().AddColumn(&Movie{}, "Year")
		},
		Rollback: func(tx *gorm.DB) error {
			type Movie struct {
				Year int
			}
			return tx.Migrator().DropColumn(&Movie{}, "Year")
		},
	}
}
