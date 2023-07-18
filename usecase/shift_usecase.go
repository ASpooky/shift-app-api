package usecase

import (
	"shift-app-go-api/model"
	"shift-app-go-api/repository"
	"shift-app-go-api/validator"
)

type IShiftUsecase interface {
	GetAllShifts(userId uint) ([]model.ShiftResponse, error)
	GetShiftById(userId uint, shiftId uint) (model.ShiftResponse, error)
	CreateShift(shift model.Shift) (model.ShiftResponse, error)
	UpdateShift(shift model.Shift, userId uint, shiftId uint) (model.ShiftResponse, error)
	DeleteShift(userId uint, shiftId uint) error
}

type shiftUsecase struct {
	sr repository.IShiftRepository
	sv validator.IShiftValidator
}

func NewShiftUsecase(sr repository.IShiftRepository, sv validator.IShiftValidator) IShiftUsecase {
	return &shiftUsecase{sr, sv}
}

func (su *shiftUsecase) GetAllShifts(userId uint) ([]model.ShiftResponse, error) {
	shifts := []model.Shift{}
	if err := su.sr.GetAllShifts(&shifts, userId); err != nil {
		return nil, err
	}
	resShifts := []model.ShiftResponse{}
	for _, v := range shifts {
		s := model.ShiftResponse{
			ID:        v.ID,
			StartTime: v.StartTime,
			EndTime:   v.StartTime,
			CreatedAt: v.EndTime,
			UpdateAt:  v.UpdateAt,
		}
		resShifts = append(resShifts, s)
	}
	return resShifts, nil
}

func (su *shiftUsecase) GetShiftById(userId uint, shiftId uint) (model.ShiftResponse, error) {
	shift := model.Shift{}
	if err := su.sr.GetShiftById(&shift, userId, shiftId); err != nil {
		return model.ShiftResponse{}, err
	}
	resShift := model.ShiftResponse{
		ID:        shift.ID,
		StartTime: shift.StartTime,
		EndTime:   shift.StartTime,
		CreatedAt: shift.EndTime,
		UpdateAt:  shift.UpdateAt,
	}
	return resShift, nil
}

func (su *shiftUsecase) CreateShift(shift model.Shift) (model.ShiftResponse, error) {
	if err := su.sv.ShiftValidate(shift); err != nil {
		return model.ShiftResponse{}, err
	}
	if err := su.sr.CreateShift(&shift); err != nil {
		return model.ShiftResponse{}, err
	}
	resShift := model.ShiftResponse{
		ID:        shift.ID,
		StartTime: shift.StartTime,
		EndTime:   shift.StartTime,
		CreatedAt: shift.EndTime,
		UpdateAt:  shift.UpdateAt,
	}
	return resShift, nil
}

func (su *shiftUsecase) UpdateShift(shift model.Shift, userId uint, shiftId uint) (model.ShiftResponse, error) {
	if err := su.sv.ShiftValidate(shift); err != nil {
		return model.ShiftResponse{}, err
	}
	if err := su.sr.UpdateShift(&shift, userId, shiftId); err != nil {
		return model.ShiftResponse{}, err
	}
	resShift := model.ShiftResponse{
		ID:        shift.ID,
		StartTime: shift.StartTime,
		EndTime:   shift.StartTime,
		CreatedAt: shift.EndTime,
		UpdateAt:  shift.UpdateAt,
	}
	return resShift, nil
}

func (su *shiftUsecase) DeleteShift(userId uint, shiftId uint) error {
	if err := su.sr.DeleteShift(userId, shiftId); err != nil {
		return err
	}
	return nil
}
