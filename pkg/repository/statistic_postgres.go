package repository

import (
	"fitter-srv/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type StatisticPostgres struct {
	db *sqlx.DB
}

func NewStatisticPostgres(db *sqlx.DB) *StatisticPostgres {
	return &StatisticPostgres{db: db}
}

func (r *StatisticPostgres) Create(userId, setId int, statistic model.Statistic) (int, error) {
	var statisticId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (set_id, user_id, date, quantity, weight, description) values ($1, $2, $3, $4, $5, $6) RETURNING id", statisticTable)
	row := r.db.QueryRow(createItemQuery, setId, userId, statistic.Date, statistic.Quantity, statistic.Weight, statistic.Description)
	if err := row.Scan(&statisticId); err != nil {
		return 0, err
	}
	return statisticId, nil
}

func (r *StatisticPostgres) GetAll(userId, setId int) ([]model.Statistic, error) {
	var statistics []model.Statistic
	query := fmt.Sprintf(`SELECT * FROM %s ts WHERE ts.set_id=$1 AND ts.user_id=$2`,
		statisticTable)
	if err := r.db.Select(&statistics, query, setId, userId); err != nil {
		return nil, err
	}
	return statistics, nil
}

func (r *StatisticPostgres) GetById(userId int, statisticId int) (model.Statistic, error) {
	var statistic model.Statistic
	query := fmt.Sprintf(`SELECT * FROM %s ts WHERE ts.id=$1 AND ts.user_id=$2`,
		statisticTable)
	if err := r.db.Get(&statistic, query, statisticId, userId); err != nil {
		return statistic, err
	}
	return statistic, nil
}

func (r *StatisticPostgres) Delete(userId int, statisticId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ts WHERE ts.id=$1 AND ts.user_id=$2`,
		statisticTable)
	_, err := r.db.Exec(query, statisticId, userId)
	return err
}

func (r *StatisticPostgres) Update(userId int, statisticId int, input model.UpdateStatisticInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Quantity != nil {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d", argId))
		args = append(args, *input.Quantity)
		argId++
	}

	if input.Weight != nil {
		setValues = append(setValues, fmt.Sprintf("weight=$%d", argId))
		args = append(args, *input.Weight)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s ts SET %s WHERE ts.id=$%d AND ts.user_id=$%d`,
		setTable, setQuery, argId, argId+1)
	args = append(args, statisticId, userId)

	_, err := r.db.Exec(query, args...)

	return err
}
