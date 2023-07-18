package validator

import (
	"shift-app-go-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IShiftValidator interface {
	ShiftValidate(shift model.Shift) error
}

type shiftValidator struct{}

func NewShiftValidator() IShiftValidator {
	return &shiftValidator{}
}

func (sv *shiftValidator) ShiftValidate(shift model.Shift) error {
	return validation.ValidateStruct(&shift,
		validation.Field(
			&shift.StartTime,
			validation.Required.Error("StartTime is required"),
		),
		validation.Field(
			&shift.EndTime,
			validation.Required.Error("EndTime is required"),
		))
}
