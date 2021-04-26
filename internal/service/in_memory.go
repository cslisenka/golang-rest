package service

import (
	"errors"

	"slisenko.com/kslisenko/golang-rest/internal/model"
)

// Service layer
type InMemoryTaskRepository struct {
	data []model.Task
}

func (tasks *InMemoryTaskRepository) InitSampleData() {
	tasks.AddTask(&model.Task{ID: "1", Description: "Task 1"})
	tasks.AddTask(&model.Task{ID: "2", Description: "Task 2"})
}

func (tasks *InMemoryTaskRepository) GetTasks() []model.Task {
	return tasks.data
}

func (tasks *InMemoryTaskRepository) AddTask(task *model.Task) {
	tasks.data = append(tasks.data, *task)
}

func (tasks *InMemoryTaskRepository) GetTaskById(id string) (model.Task, error) {
	for _, task := range tasks.data {
		if task.ID == id {
			return task, nil
		}
	}

	return model.Task{}, errors.New("task not found")
}
