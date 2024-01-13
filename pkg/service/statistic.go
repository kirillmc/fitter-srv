package service

import (
	"fitter-srv/model"
	"fitter-srv/pkg/repository"
)

type StatisticService struct {
	repo    repository.Statistic
	setRepo repository.Set
}

func NewStatisticService(repo repository.Statistic, setRepo repository.Set) *StatisticService {
	return &StatisticService{repo: repo, setRepo: setRepo}
}

func (s *StatisticService) Create(userId, setId int, statistic model.Statistic) (int, error) {
	_, err := s.setRepo.GetById(userId, setId)
	if err != nil {
		//list does not exist or does not belong to user
		return 0, err
	}

	return s.repo.Create(userId, setId, statistic)
}

func (s *StatisticService) GetAll(userId, setId int) ([]model.Statistic, error) {
	return s.repo.GetAll(userId, setId)
}

func (s *StatisticService) GetById(userId, statisticId int) (model.Statistic, error) {
	return s.repo.GetById(userId, statisticId)
}

func (s *StatisticService) Delete(userId, statisticId int) error {
	return s.repo.Delete(userId, statisticId)
}

func (s *StatisticService) Update(userId int, statisticId int, input model.UpdateStatisticInput) error {
	return s.repo.Update(userId, statisticId, input)
}
