package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	generatedModel "task-managament/graph/model"
	"task-managament/server/commons"
)

type databaseRepository struct {
	con *sql.DB
}

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
		return nil, errors.New("Failed while migration")
	}
	return &databaseRepository{con}, nil
}

func (repo *databaseRepository) Save(ctx context.Context, model interface{}) (data interface{}, err error) {
	taskInput, _ := model.(*generatedModel.TaskInput)
	//TODO: update below print statement with database persist functionality
	log.Panic(ctx, "type cast model to eoSkeleton %v", taskInput)
	return nil, err
}

func (repo *databaseRepository) Update(ctx context.Context, model interface{}) (data interface{}, err error) {
	//TODO: update below print statement with database persist functionality
	return nil, err
}

func (repo *databaseRepository) List(ctx context.Context, id int) (model []interface{}, err error) {
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
