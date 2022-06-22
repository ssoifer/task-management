package repositories

import (
	"context"
	"errors"
	"task-managament/server/commons"
	"task-managament/server/repositories/db"
)

type Repository interface {
	Save(ctx context.Context, model interface{}) (data interface{}, err error)
	Update(ctx context.Context, model interface{}) (data interface{}, err error)
	List(ctx context.Context, id int) (model []interface{}, err error)
}

func GetRepository(repoType commons.RepositoryType) (Repository, error) {
	if repoType == commons.RepositoryTypeDB {
		return db.NewRepository()
	} else {
		return nil, errors.New("Invalid Repository Type provided")
	}
}
