package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"task-managament/graph/generated"
	"task-managament/graph/model"
	"time"

	"github.com/dgryski/trifles/uuid"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, input model.TaskInput) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTasks(ctx context.Context) ([]*model.Task, error) {
	//panic(fmt.Errorf("not implemented"))
	var tasks []*model.Task
	dummyTask := model.Task{
		Title:     "our dummy link",
		ID:        uuid.UUIDv4(),
		Content:   "Hello World",
		Views:     1,
		Timestamp: time.Now().String(),
	}
	tasks = append(tasks, &dummyTask)
	return tasks, nil
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
