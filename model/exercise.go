package model

import "errors"

type Exercise struct {
	Id              int    `json:"id" db:"id"`
	DayId           int    `json:"day_id" db:"day_id"`
	ExerciseName    string `json:"exercise_name" db:"exercise_name" binding:"required"`
	ExercisePicture string `json:"exercise_picture" db:"exercise_picture"`
	Description     string `json:"description" db:"description"`
}

// Для обновления стурктура с указателями
type UpdateExerciseInput struct {
	ExerciseName    *string `json:"exercise_name"`
	ExercisePicture *string `json:"exercise_picture"`
	Description     *string `json:"description"`
}

func (i UpdateExerciseInput) Validate() error {
	if i.ExerciseName == nil && i.ExercisePicture == nil && i.Description == nil {
		return errors.New("update train exercise has no values to update")
	}
	return nil
}
