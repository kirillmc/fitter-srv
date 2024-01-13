package service

import (
	"fitter-srv/model"
	"fitter-srv/pkg/repository"
)

type SetService struct {
	repo         repository.Set
	exerciseRepo repository.Exercise
}

func NewSetService(repo repository.Set, exerciseRepo repository.Exercise) *SetService {
	return &SetService{repo: repo, exerciseRepo: exerciseRepo}
}

func (s *SetService) Create(userId, exerciseId int, set model.Set) (int, error) {
	_, err := s.exerciseRepo.GetById(userId, exerciseId)
	if err != nil {
		//list does not exist or does not belong to user
		return 0, err
	}

	return s.repo.Create(exerciseId, set)
}

func (s *SetService) GetAll(userId, exerciseId int) ([]model.Set, error) {
	return s.repo.GetAll(userId, exerciseId)
}

func (s *SetService) GetById(userId, setId int) (model.Set, error) {
	return s.repo.GetById(userId, setId)
}

func (s *SetService) Delete(userId, setId int) error {
	return s.repo.Delete(userId, setId)
}

func (s *SetService) Update(userId int, setId int, input model.UpdateSetInput) error {
	return s.repo.Update(userId, setId, input)
}
