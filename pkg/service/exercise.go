package service

import (
	"fitter-srv/model"
	"fitter-srv/pkg/repository"
)

type ExerciseService struct {
	repo    repository.Exercise
	dayRepo repository.Day
}

func NewExerciseService(repo repository.Exercise, dayRepo repository.Day) *ExerciseService {
	return &ExerciseService{repo: repo, dayRepo: dayRepo}
}

func (s *ExerciseService) Create(userId, dayId int, exercise model.Exercise) (int, error) {
	_, err := s.dayRepo.GetById(userId, dayId)
	if err != nil {
		//list does not exist or does not belong to user
		return 0, err
	}

	return s.repo.Create(dayId, exercise)
}

func (s *ExerciseService) GetAll(userId, dayId int) ([]model.Exercise, error) {
	return s.repo.GetAll(userId, dayId)
}

func (s *ExerciseService) GetById(userId, exerciseId int) (model.Exercise, error) {
	return s.repo.GetById(userId, exerciseId)
}

func (s *ExerciseService) Delete(userId, exerciseId int) error {
	return s.repo.Delete(userId, exerciseId)
}

func (s *ExerciseService) Update(userId int, exerciseId int, input model.UpdateExerciseInput) error {
	return s.repo.Update(userId, exerciseId, input)
}
