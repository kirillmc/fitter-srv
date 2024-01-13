package model

import "errors"

type Statistic struct {
	Id          int     `json:"id" db:"id"`
	SetId       int     `json:"set_id" db:"set_id"`
	UserId      int     `json:"user_id" db:"user_id"`
	Date        string  `json:"date" db:"date"`
	Quantity    int     `json:"quantity" db:"quantity"`
	Weight      float64 `json:"weight" db:"weight"`
	Description string  `json:"description" db:"description"`
}

// Для обновления стурктура с указателями
type UpdateStatisticInput struct {
	Quantity    *int     `json:"quantity"`
	Weight      *float64 `json:"weight"`
	Description *string  `json:"description"`
}

func (i UpdateStatisticInput) Validate() error {
	if i.Quantity == nil && i.Weight == nil && i.Description == nil {
		return errors.New("update train statistic has no values to update")
	}
	return nil
}
