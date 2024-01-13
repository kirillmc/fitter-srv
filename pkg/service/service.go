package service

import (
	"fitter-srv/model"
	"fitter-srv/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Program interface {
	Create(userId int, program model.TrainProgram) (int, error)
	GetAll(userId int) ([]model.TrainProgram, error)
	GetById(userId, programId int) (model.TrainProgram, error)
	Update(userId, id int, input model.UpdateProgramInput) error
	Delete(userId, programId int) error
}

type Day interface {
	Create(userId, programId int, day model.TrainDay) (int, error)
	GetAll(userId, programId int) ([]model.TrainDay, error)
	GetById(userId, dayId int) (model.TrainDay, error)
	Update(userId int, dayId int, input model.UpdateDayInput) error
	Delete(userId, dayId int) error
}
type Exercise interface {
	Create(userId, dayId int, exercise model.Exercise) (int, error)
	GetAll(userId, dayId int) ([]model.Exercise, error)
	GetById(userId, Set int) (model.Exercise, error)
	Update(userId int, erciseId int, input model.UpdateExerciseInput) error
	Delete(userId, erciseId int) error
}
type Set interface {
	Create(userId, exerciseId int, set model.Set) (int, error)
	GetAll(userId, exerciseId int) ([]model.Set, error)
	GetById(userId, setId int) (model.Set, error)
	Update(userId int, setId int, input model.UpdateSetInput) error
	Delete(userId, setId int) error
}
type Statistic interface {
	Create(userId, setId int, statistic model.Statistic) (int, error)
	GetAll(userId, setId int) ([]model.Statistic, error)
	GetById(userId, statisticId int) (model.Statistic, error)
	Update(userId int, statisticId int, input model.UpdateStatisticInput) error
	Delete(userId, statisticId int) error
}

type Service struct {
	Authorization
	Program
	Day
	Exercise
	Set
	Statistic
}

// Указатель на структуру репозитория - внедрение зависимостей(Dependency injection)
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Program:       NewProgramService(repos.Program),
		Day:           NewDayService(repos.Day, repos.Program),
		Exercise:      NewExerciseService(repos.Exercise, repos.Day),
		Set:           NewSetService(repos.Set, repos.Exercise),
		Statistic:     NewStatisticService(repos.Statistic, repos.Set),
	}
}
