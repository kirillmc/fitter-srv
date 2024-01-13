package model

import "errors"

type TrainProgram struct {
	Id          int    `json:"id" db:"id"`
	UserId      int    `json:"user_id" db:"user_id"`
	ProgramName string `json:"program_name" db:"program_name" binding:"required"`
	Description string `json:"description" db:"description"`
	IsPublic    bool   `json:"is_public" db:"is_public"`
}

// Для обновления стурктура с указателями
type UpdateProgramInput struct {
	ProgramName *string `json:"program_name"`
	Description *string `json:"description"`
	IsPublic    *bool   `json:"is_public"`
}

func (i UpdateProgramInput) Validate() error {
	if i.ProgramName == nil && i.Description == nil && i.IsPublic == nil {
		return errors.New("update train program has no values to update")
	}
	return nil
}
