package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/olusamimaths/AnonymousBoard/models"
	"gorm.io/gorm"
)

func synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:       "initial",
			Migrate:  migrate,
			Rollback: rollback,
		},
	})
	return m.Migrate()
}

func migrate(tx *gorm.DB) error {
	if err := tx.AutoMigrate(
		&models.Thread{},
		&models.Reply{},
	); err != nil {
		return err
	}

	if err := tx.Exec("ALTER TABLE threads ADD CONSTRAINT fk_ FOREIGN KEY (thread_id) REFERENCES thread (id)").Error; err != nil {
		return err
	}
	return nil
}

func rollback(tx *gorm.DB) error {
	return tx.Migrator().DropTable(
		&models.Thread{},
		&models.Reply{},
	)
}
