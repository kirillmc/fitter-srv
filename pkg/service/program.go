package service

import (
	"fitter-srv/model"
	"fitter-srv/pkg/repository"
)

type ProgramService struct {
	repo repository.Program
}

func (p ProgramService) Create(userId int, program model.TrainProgram) (int, error) {
	return p.repo.Create(userId, program)
}

func (p ProgramService) GetAll(userId int) ([]model.TrainProgram, error) {
	return p.repo.GetAll(userId)
}

func (p ProgramService) GetById(userId, programId int) (model.TrainProgram, error) {
	return p.repo.GetById(userId, programId)
}

func (p ProgramService) Update(userId, id int, input model.UpdateProgramInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return p.repo.Update(userId, id, input)
}

func (p ProgramService) Delete(userId, programId int) error {
	return p.repo.Delete(userId, programId)
}

func NewProgramService(repo repository.Program) *ProgramService { return &ProgramService{repo: repo} }
