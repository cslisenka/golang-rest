package service

import "slisenko.com/kslisenko/golang-rest/internal/model"

type TaskService interface {
	InitSampleData()

	GetTasks() []model.Task

	AddTask(task *model.Task)

	GetTaskById(id string) (model.Task, error)
}
