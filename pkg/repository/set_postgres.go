package repository

import (
	"fitter-srv/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SetPostgres struct {
	db *sqlx.DB
}

func NewSetPostgres(db *sqlx.DB) *SetPostgres {
	return &SetPostgres{db: db}
}

func (r *SetPostgres) Create(exerciseId int, set model.Set) (int, error) {
	var setId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (exercise_id, quantity, weight, description) values ($1, $2, $3, $4) RETURNING id", setTable)
	row := r.db.QueryRow(createItemQuery, exerciseId, set.Quantity, set.Weight, set.Description)
	if err := row.Scan(&setId); err != nil {
		return 0, err
	}
	return setId, nil
}

func (r *SetPostgres) GetAll(userId, exerciseId int) ([]model.Set, error) {
	var sets []model.Set
	query := fmt.Sprintf(`SELECT * FROM %s ts WHERE ts.exercise_id=$1`,
		setTable)
	if err := r.db.Select(&sets, query, exerciseId); err != nil {
		return nil, err
	}
	return sets, nil
}

func (r *SetPostgres) GetById(userId int, setId int) (model.Set, error) {
	var set model.Set
	query := fmt.Sprintf(`SELECT * FROM %s ts WHERE ts.id=$1`,
		setTable)
	if err := r.db.Get(&set, query, setId); err != nil {
		return set, err
	}
	return set, nil
}

func (r *SetPostgres) Delete(userId int, setId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ts WHERE ts.id=$1`,
		setTable)
	_, err := r.db.Exec(query, setId)
	return err
}

func (r *SetPostgres) Update(userId int, setId int, input model.UpdateSetInput) error {
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
	query := fmt.Sprintf(`UPDATE %s ts SET %s WHERE ts.id=$%d`,
		setTable, setQuery, argId)
	args = append(args, setId)

	_, err := r.db.Exec(query, args...)

	return err
}
