package service

import (
	"fitter-srv/model"
	"fitter-srv/pkg/repository"
)

type DayService struct {
	repo        repository.Day
	programRepo repository.Program
}

func NewDayService(repo repository.Day, programRepo repository.Program) *DayService {
	return &DayService{repo: repo, programRepo: programRepo}
}

func (s *DayService) Create(userId, programId int, day model.TrainDay) (int, error) {
	_, err := s.programRepo.GetById(userId, programId)
	if err != nil {
		//list does not exist or does not belong to user
		return 0, err
	}

	return s.repo.Create(programId, day)
}

func (s *DayService) GetAll(userId, programId int) ([]model.TrainDay, error) {
	return s.repo.GetAll(userId, programId)
}

func (s *DayService) GetById(userId, dayId int) (model.TrainDay, error) {
	return s.repo.GetById(userId, dayId)
}

func (s *DayService) Delete(userId, dayId int) error {
	return s.repo.Delete(userId, dayId)
}

func (s *DayService) Update(userId int, dayId int, input model.UpdateDayInput) error {
	return s.repo.Update(userId, dayId, input)
}
