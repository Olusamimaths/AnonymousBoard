package services

import (
	"errors"

	"github.com/olusamimaths/AnonymousBoard/database"
	"github.com/olusamimaths/AnonymousBoard/models"
	"github.com/olusamimaths/AnonymousBoard/utils"
	"gorm.io/gorm"
)

type ThreadService interface {
	List(page int) (error, *[]models.Thread)
	GetByID(id string) (error, *models.Thread)
	Create(t models.Thread) (error, *models.Thread)
	Report(id string) error
	DeleteWithPassword(id, password string) error
}

type threadService struct {
	db *gorm.DB
}

func (ts *threadService) List(page int) (error, *[]models.Thread) {
	var t []models.Thread
	offset := 0

	if page > 0 {
		offset = page - 1
	}

	results := ts.db.
		Preload("Replies", func (db *gorm.DB) *gorm.DB {
			return db.Limit(10)
		}).
		Order("bumped_on DESC").
		Limit(10).
		Offset(offset).
		Find(&t)
	return results.Error, &t
}

func (ts *threadService) GetByID(id string) (error, *models.Thread) {
	var t models.Thread
	result := ts.db.
		Where("id = ?", id).
		Preload("Replies", func (db *gorm.DB) *gorm.DB {
			return db.Limit(10)
		}).
		First(&t)

	return result.Error, &t
}

func (ts *threadService) Create(t models.Thread) (error, *models.Thread) {
	hashedPassword, err := utils.HashPassword(t.DeletePassword)

	if err != nil {
		return err, nil
	}
	t.DeletePassword = hashedPassword

	result := ts.db.Create(&t)

	return result.Error, &t
}

func (ts *threadService) Report(id string) error {
	return ts.db.Transaction(func(tx *gorm.DB) error {
		var t models.Thread
		if result := tx.Where("id = ?", id).First(&t); result.Error != nil {
			return result.Error
		}

		if result := tx.Model(&t).Where("id = ?", id).Update("reported", true); result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func (ts *threadService) DeleteWithPassword(id, password string) error {
	return ts.db.Transaction(func(tx *gorm.DB) error {
		var t models.Thread
		if result := tx.Where("id = ?", id).First(&t); result.Error != nil {
			return result.Error
		}

		if !utils.IsPasswordValid(t.DeletePassword, password) {
			return errors.New("incorrect password")
		}

		if result := tx.Model(&t).Where("id = ?", id).Update("text", "[deleted]"); result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func NewThreadService(conn database.DatabaseConnection) ThreadService {
	return &threadService{db: conn.Get()}
}
