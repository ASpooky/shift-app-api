package repository

import (
	"fmt"
	"shift-app-go-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IShiftRepository interface {
	GetAllShifts(shift *[]model.Shift, userId uint) error
	GetShiftById(shift *model.Shift, userId uint, shiftId uint) error
	CreateShift(shift *model.Shift) error
	UpdateShift(shift *model.Shift, userId uint, shiftId uint) error
	DeleteShift(userId uint, shiftId uint) error
}

type shiftRepository struct {
	db *gorm.DB
}

func NewShiftRepository(db *gorm.DB) IShiftRepository {
	return &shiftRepository{db}
}

func (sr *shiftRepository) GetAllShifts(shifts *[]model.Shift, userId uint) error {
	if err := sr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(shifts).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shiftRepository) GetShiftById(shift *model.Shift, userId uint, shiftId uint) error {
	if err := sr.db.Create(shift).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shiftRepository) CreateShift(shift *model.Shift) error {
	if err := sr.db.Create(shift).Error; err != nil {
		return err
	}
	return nil
}

func (sr *shiftRepository) UpdateShift(shift *model.Shift, userId uint, shiftId uint) error {
	result := sr.db.Model(shift).Clauses(clause.Returning{}).Where("id=? AND user_id=?", shiftId, userId).Updates(shift)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (sr *shiftRepository) DeleteShift(userId uint, shiftId uint) error {
	result := sr.db.Where("id=? AND user_id=?", shiftId, userId).Delete(&model.Shift{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
