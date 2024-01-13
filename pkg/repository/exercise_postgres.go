package repository

import (
	"fitter-srv/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ExercisePostgres struct {
	db *sqlx.DB
}

func NewExercisePostgres(db *sqlx.DB) *ExercisePostgres {
	return &ExercisePostgres{db: db}
}

func (r *ExercisePostgres) Create(dayId int, exercise model.Exercise) (int, error) {
	var exerciseId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (day_id, exercise_name, exercise_picture, description) values ($1, $2, $3, $4) RETURNING id", exerciseTable)
	row := r.db.QueryRow(createItemQuery, dayId, exercise.ExerciseName, exercise.ExercisePicture, exercise.Description)
	if err := row.Scan(&exerciseId); err != nil {
		return 0, err
	}
	return exerciseId, nil
}

func (r *ExercisePostgres) GetAll(userId, dayId int) ([]model.Exercise, error) {
	var exercises []model.Exercise
	query := fmt.Sprintf(`SELECT * FROM %s te WHERE te.day_id=$1`,
		exerciseTable)
	if err := r.db.Select(&exercises, query, dayId); err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *ExercisePostgres) GetById(userId int, exerciseId int) (model.Exercise, error) {
	var exercise model.Exercise
	query := fmt.Sprintf(`SELECT * FROM %s te WHERE te.id=$1`,
		exerciseTable)
	if err := r.db.Get(&exercise, query, exerciseId); err != nil {
		return exercise, err
	}
	return exercise, nil
}

func (r *ExercisePostgres) Delete(userId int, exerciseId int) error {
	query := fmt.Sprintf(`DELETE FROM %s te WHERE te.id=$1`,
		exerciseTable)
	_, err := r.db.Exec(query, exerciseId)
	return err
}

func (r *ExercisePostgres) Update(userId int, exerciseId int, input model.UpdateExerciseInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ExerciseName != nil {
		setValues = append(setValues, fmt.Sprintf("exercise_name=$%d", argId))
		args = append(args, *input.ExerciseName)
		argId++
	}

	if input.ExercisePicture != nil {
		setValues = append(setValues, fmt.Sprintf("exercise_picture=$%d", argId))
		args = append(args, *input.ExercisePicture)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s te SET %s WHERE te.id=$%d`,
		exerciseTable, setQuery, argId)
	args = append(args, exerciseId)

	_, err := r.db.Exec(query, args...)

	return err
}
