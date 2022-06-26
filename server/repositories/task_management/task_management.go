package taskmanagement

import (
	"context"
	"errors"
	"task-managament/server/api/graph/model"
	"task-managament/server/commons"
	dbModel "task-managament/server/models"
	"task-managament/server/repositories/db"
)

type TaskRepository interface {
	Save(id string, content string, title string, views int, timestamp string) (dbModel *dbModel.Task, err error)
	Update(ctx context.Context, input model.TaskInput) (model *model.Task, err error)
	GetList() (dbModel []dbModel.Task, err error)
	GetById(ctx context.Context, id int) (dbModel *dbModel.Task, err error)
}

func GetRepository(repoType commons.RepositoryType) (TaskRepository, error) {
	if repoType == commons.RepositoryTypeDB {
		return db.NewRepository()
	} else {
		return nil, errors.New("Invalid Repository Type provided")
	}
}
