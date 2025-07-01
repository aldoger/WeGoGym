package migration

import (
	"fmt"

	"go-kpl/internal/domain/models"
	mylog "go-kpl/internal/pkg/logger"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println(mylog.ColorizeInfo("\n=========== Start Migrate ==========="))
	mylog.Infof("Migrating Tables...")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Membership{},
		&models.UserMembership{},
		&models.EntryHistory{},
		&models.UserPersonalTrainer{},
	); err != nil {
		return err
	}

	return nil
}
