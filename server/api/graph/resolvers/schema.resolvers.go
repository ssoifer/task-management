package resolvers

import (
	"context"
	"fmt"
	"github.com/dgryski/trifles/uuid"
	"log"
	"task-managament/server/api/graph/generated"
	"task-managament/server/api/graph/model"
	"task-managament/server/commons"
)

//This file will be automatically regenerated based on the schema, any resolver implementations
//will be copied through when generating and any unknown code will be moved to the end.

func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	uuidString := uuid.UUIDv4()
	res, err := r.TasksManagementService.CreateTask(uuidString, input.Content, input.Title, input.Views, input.Timestamp)
	if err != nil {
		log.Panic(ctx, err, "")
		if err.Code == commons.ErrorCodeBadUserInput {
			return nil, getGqlError(ctx, "Unable to create task - duplicate service tag", err.Code, err.Description)
		}
		return nil, getGqlError(ctx, "Unable to create ece inventory", err.Code, err.Description)
	}
	return mapDbToModel(res), nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.TaskInput) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTasks(ctx context.Context) ([]*model.Task, error) {

	var list []*model.Task

	res, err := r.TasksManagementService.GetTaskList()
	for idx := range res {
		list = append(list, mapDbToModel(&res[idx]))
	}
	if err != nil {
		log.Print(ctx, err, "")
		return list, getGqlError(ctx, "Unable to get list of tasks", err.Code, err.Description)
	}
	return list, nil
}

func (r *queryResolver) GetTaskByID(ctx context.Context, id string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
