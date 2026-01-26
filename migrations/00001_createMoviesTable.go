package migrations

import (
	"go-gin-mysql/models"

	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateMoviesTableMigration() *gormigrate.Migration {

	return &gormigrate.Migration{
		ID: "00001_createMoviesTable",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Movie{})

		},

		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("movies")
		},
	}
}
