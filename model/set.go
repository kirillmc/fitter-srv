package model

import "errors"

type Set struct {
	Id          int     `json:"id" db:"id"`
	ExerciseId  int     `json:"exercise_id" db:"exercise_id"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Weight      float64 `json:"weight" db:"weight"`
	Description string  `json:"description" db:"description"`
}

// Для обновления стурктура с указателями
type UpdateSetInput struct {
	Quantity    *int     `json:"quantity"`
	Weight      *float64 `json:"weight"`
	Description *string  `json:"description"`
}

func (i UpdateSetInput) Validate() error {
	if i.Quantity == nil && i.Weight == nil && i.Description == nil {
		return errors.New("update train sets has no values to update")
	}
	return nil
}
