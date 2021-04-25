package service

import (
	"errors"

	"slisenko.com/kslisenko/golang-rest/internal/model"
)

// Service layer
// TODO change to map
var tasks []model.Task

func InitSampleData() {
	AddTask(&model.Task{ID: "1", Description: "Task 1"})
	AddTask(&model.Task{ID: "2", Description: "Task 2"})
}

func GetTasks() []model.Task {
	return tasks
}

func AddTask(task *model.Task) {
	tasks = append(tasks, *task)
}

func GetTaskById(id string) (model.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return model.Task{}, errors.New("task not found")
}
