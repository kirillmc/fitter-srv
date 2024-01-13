package repository

import (
	"fitter-srv/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DayPostgres struct {
	db *sqlx.DB
}

func NewDayPostgres(db *sqlx.DB) *DayPostgres {
	return &DayPostgres{db: db}
}

func (r *DayPostgres) Create(programId int, day model.TrainDay) (int, error) {
	var dayId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (program_id, day_name, description) values ($1, $2, $3) RETURNING id", dayTable)
	row := r.db.QueryRow(createItemQuery, programId, day.DayName, day.Description)
	if err := row.Scan(&dayId); err != nil {
		return 0, err
	}
	return dayId, nil
}

func (r *DayPostgres) GetAll(userId, programId int) ([]model.TrainDay, error) {
	var days []model.TrainDay
	query := fmt.Sprintf(`SELECT * FROM %s td WHERE td.program_id=$1`,
		dayTable)
	if err := r.db.Select(&days, query, programId); err != nil {
		return nil, err
	}
	return days, nil
}

func (r *DayPostgres) GetById(userId int, dayId int) (model.TrainDay, error) {
	var day model.TrainDay
	query := fmt.Sprintf(`SELECT * FROM %s td WHERE td.id=$1`,
		dayTable)
	if err := r.db.Get(&day, query, dayId); err != nil {
		return day, err
	}
	return day, nil
}

func (r *DayPostgres) Delete(userId int, dayId int) error {
	query := fmt.Sprintf(`DELETE FROM %s td WHERE td.id=$1`,
		dayTable)
	_, err := r.db.Exec(query, dayId)
	return err
}

func (r *DayPostgres) Update(userId int, dayId int, input model.UpdateDayInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.DayName != nil {
		setValues = append(setValues, fmt.Sprintf("day_name=$%d", argId))
		args = append(args, *input.DayName)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s td SET %s WHERE td.id=$%d`,
		dayTable, setQuery, argId)
	args = append(args, dayId)

	_, err := r.db.Exec(query, args...)

	return err
}
