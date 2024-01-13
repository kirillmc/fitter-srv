package repository

import (
	"fitter-srv/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type ProgramPostgres struct {
	db *sqlx.DB
}

func NewProgramPostgres(db *sqlx.DB) *ProgramPostgres {
	return &ProgramPostgres{db: db}
}

func (r *ProgramPostgres) Create(userId int, program model.TrainProgram) (int, error) {
	var id int
	createProgramQuery := fmt.Sprintf("INSERT INTO %s (user_id, program_name, description, is_public) VALUES ($1, $2, $3, $4) RETURNING id", programTable)
	row := r.db.QueryRow(createProgramQuery, userId, program.ProgramName, program.Description, program.IsPublic)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProgramPostgres) GetAll(userId int) ([]model.TrainProgram, error) {
	var programs []model.TrainProgram
	query := fmt.Sprintf("SELECT * FROM %s tp WHERE tp.user_id = $1",
		programTable)
	err := r.db.Select(&programs, query, userId)
	return programs, err
}

func (r *ProgramPostgres) GetById(userId, programId int) (model.TrainProgram, error) {
	var program model.TrainProgram
	query := fmt.Sprintf("SELECT * FROM %s tp WHERE tp.user_id = $1 AND tp.id=$2",
		programTable)
	err := r.db.Get(&program, query, userId, programId)
	return program, err
}

func (r *ProgramPostgres) Delete(userId, programId int) error {
	query := fmt.Sprintf("DELETE FROM %s tp WHERE tp.user_id=$1 AND tp.id=$2", programTable)

	/*
		В некоторых базах данных SQL (например, в PostgreSQL) USING можно использовать
		в операторах UPDATE и DELETE вместе с псевдонимами таблиц.
		Это может быть особенно полезно, когда вы хотите обновить или удалить записи на основе сравнения с другой таблицей.
	*/

	_, err := r.db.Exec(query, userId, programId)

	return err
}

func (r *ProgramPostgres) Update(userId, programId int, input model.UpdateProgramInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ProgramName != nil {
		setValues = append(setValues, fmt.Sprintf("program_name=$%d", argId))
		args = append(args, *input.ProgramName)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.IsPublic != nil {
		setValues = append(setValues, fmt.Sprintf("is_public=$%d", argId))
		args = append(args, *input.IsPublic)
		argId++
	}
	// title = 'updated title'
	// description = 'updated description'
	// title = 'updated title' description = 'updated description'
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tp SET %s WHERE tp.id=$%d AND tp.user_id=$%d",
		programTable, setQuery, argId, argId+1)
	args = append(args, programId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
