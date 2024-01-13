package model

import "errors"

type TrainDay struct {
	Id          int    `json:"id" db:"id"`
	ProgramId   int    `json:"program_id" db:"program_id"`
	DayName     string `json:"day_name" db:"day_name" binding:"required"`
	Description string `json:"description" db:"description"`
}

// Для обновления стурктура с указателями
type UpdateDayInput struct {
	DayName     *string `json:"day_name"`
	Description *string `json:"description"`
}

func (i UpdateDayInput) Validate() error {
	if i.DayName == nil && i.Description == nil {
		return errors.New("update train day has no values to update")
	}
	return nil
}
