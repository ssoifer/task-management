package resolvers

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	gqlModel "task-managament/server/api/graph/model"
	"task-managament/server/commons"
	dbModel "task-managament/server/models"
)

func getGqlError(ctx context.Context, message string, errCode commons.ErrorCode, errDescription commons.ErrorDescription) error {
	return &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: message,
		Extensions: map[string]interface{}{
			"code":        errCode,
			"description": errDescription,
		},
	}
}

func mapDbToModel(dbModel *dbModel.Task) *gqlModel.Task {
	if dbModel == nil {
		return nil
	}

	return &gqlModel.Task{
		ID:        dbModel.ID,
		Title:     dbModel.Title,
		Content:   dbModel.Content,
		Views:     dbModel.Views,
		Timestamp: dbModel.Timestamp,
	}
}
