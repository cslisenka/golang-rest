package service

import "slisenko.com/kslisenko/golang-rest/internal/model"

type TaskService interface {
	Init() error

	InitSampleData() error

	GetTasks() ([]model.Task, error)

	AddTask(task *model.Task) error

	GetTaskById(id string) (model.Task, error)
}
