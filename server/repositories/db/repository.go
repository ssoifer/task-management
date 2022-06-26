package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	generatedModel "task-managament/server/api/graph/model"
	"task-managament/server/commons"
	dbModel "task-managament/server/models"
)

var saveQuery = fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1,$2,$3,$4,$5) RETURNING %s;", repositoryTableName, insertFields, returnFields)
var getQuery = fmt.Sprintf("SELECT %s FROM %s WHERE service_tag = $1;", returnFields, repositoryTableName)

type databaseRepository struct {
	con *sql.DB
}

const (
	repositoryTableName = "task"
	insertFields        = " id, content, title, views, timestamp "
	returnFields        = insertFields
)

func NewRepository() (*databaseRepository, error) {
	ctx := context.Background()
	config := &Config{}
	*config = parseEnv()
	con, err := NewDatabase(*config)
	if err != nil {
		log.Panic(ctx, err, "unable to setup database")
		return nil, errors.New("unable to setup database")
	}

	err = Migrate(con)
	if err != nil {
		log.Panic(ctx, err, "unable to setup database")
		return nil, errors.New("Failed during migration")
	}
	return &databaseRepository{con}, nil
}

func (repo *databaseRepository) Save(id string, content string, title string, views int, timestamp string) (*dbModel.Task, error) {
	taskInput := dbModel.Task{}
	err := repo.con.QueryRow(saveQuery, id, content, title, views, timestamp).Scan(&taskInput.ID,
		&taskInput.Content, &taskInput.Title, &taskInput.Views, &taskInput.Timestamp)
	if err == nil {
		return &taskInput, nil
	}
	return nil, err
}

func (repo *databaseRepository) Update(ctx context.Context, input generatedModel.TaskInput) (model *generatedModel.Task, err error) {
	//TODO: update below print statement with database persist functionality
	return nil, err
}

func (repo *databaseRepository) GetList() (model []dbModel.Task, _ error) {
	baseQuery := fmt.Sprintf("SELECT %s FROM %s ", returnFields, repositoryTableName)

	var rows *sql.Rows
	var err error = nil
	listQuery := fmt.Sprintf(baseQuery)
	rows, err = repo.con.Query(listQuery)
	if err != nil {
		return []dbModel.Task{}, err
	}

	var list []dbModel.Task
	for rows.Next() {
		var taskItem dbModel.Task
		err := rows.Scan(&taskItem.ID, &taskItem.Title, &taskItem.Content, &taskItem.Views, &taskItem.Timestamp)
		if err != nil {
			return []dbModel.Task{}, err
		}
		list = append(list, taskItem)
	}
	if err = rows.Err(); err != nil {
		return []dbModel.Task{}, err
	}
	return list, nil
}

func (repo *databaseRepository) GetById(ctx context.Context, id int) (model *dbModel.Task, err error) {
	//TODO: update below print statement with database persist functionality
	return nil, err
}

func parseEnv() Config {
	return Config{
		Host:     "localhost",
		Port:     "9097",
		User:     "postgres",
		Password: "123456",
		Database: commons.DatabaseName,
	}
}
