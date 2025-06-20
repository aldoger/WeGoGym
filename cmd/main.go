package cmd

import (
	"fmt"
	"go-kpl/database"
	"os"

	"go-kpl/database/migration"
	mylog "go-kpl/internal/pkg/logger"

	"gorm.io/gorm"
)

func Commands() error {
	db := database.New()
	if err := getParams(db); err != nil {
		return err
	}
	return nil
}

func getParams(db *gorm.DB) error {
	migrate := false
	seeder := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seeder" {
			seeder = true
		}
	}

	if migrate {
		if err := migration.Migrate(db); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
		mylog.Infof("Migration completed successfully")
	}

	if seeder {
		mylog.Infof("Seeder has not been set")
	}

	if seeder || migrate {
		os.Exit(0)
	}

	return nil
}
