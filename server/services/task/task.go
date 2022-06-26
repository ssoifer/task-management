package task

import (
	"context"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"task-managament/server/api/graph/model"
	"task-managament/server/commons"
	dbModel "task-managament/server/models"
	taskmanagement "task-managament/server/repositories/task_management"
)

type TasksManagementService interface {
	GetTaskById(serviceTag string, id string) (*dbModel.Task, *commons.Error)
	GetTaskList() ([]dbModel.Task, *commons.Error)
	CreateTask(id string, content string, title string, views int, timestamp string) (*dbModel.Task, *commons.Error)
	UpdateTask(serviceTag string, input model.TaskInput) (*model.Task, error)
}

type service struct {
	taskRepository taskmanagement.TaskRepository
}

func New(taskRepository taskmanagement.TaskRepository) TasksManagementService {
	return &service{
		taskRepository: taskRepository,
	}
}

func (s service) GetTaskById(serviceTag string, id string) (*dbModel.Task, *commons.Error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetTaskList() ([]dbModel.Task, *commons.Error) {
	ctx := context.Background()
	res, err := s.taskRepository.GetList()
	if err == nil {
		log.Print(ctx, "successfully retrieved task list")
		return res, nil
	}
	return []dbModel.Task{}, commons.NewError(commons.ErrorCodeNotFound, commons.ErrorDescriptionNotFound, err)
}

func (s service) CreateTask(id string, content string, title string, views int, timestamp string) (*dbModel.Task, *commons.Error) {
	ctx := context.Background()
	res, err := s.taskRepository.Save(id, content, title, views, timestamp)
	if err == nil {
		log.Print(ctx, "created new task ", res)
		return res, nil
	}

	var pqError *pq.Error
	if errors.As(err, &pqError) {
		if pqError.Code == "23505" { // unique_violation
			return nil, commons.NewError(commons.ErrorCodeBadUserInput, commons.ErrorDescriptionBadUserInput, err)
		}
	}
	return nil, commons.NewError(commons.ErrorCodeInternalServerError, commons.ErrorDescriptionInternalServerError, err)
}

func (s service) UpdateTask(serviceTag string, input model.TaskInput) (*model.Task, error) {
	//TODO implement me
	panic("implement me")
}
