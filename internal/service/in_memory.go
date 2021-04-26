package service

import (
	"errors"

	"slisenko.com/kslisenko/golang-rest/internal/model"
)

// Service layer
type InMemoryTaskRepository struct {
	data []model.Task
}

func (tasks *InMemoryTaskRepository) Init() error {
	// Nothing to do
	return nil
}

func (tasks *InMemoryTaskRepository) InitSampleData() error {
	err1 := tasks.AddTask(&model.Task{ID: "1", Description: "Task 1"})
	err2 := tasks.AddTask(&model.Task{ID: "2", Description: "Task 2"})

	if err1 != nil || err2 != nil {
		return errors.New(err1.Error() + err2.Error())
	}

	return nil
}

func (tasks *InMemoryTaskRepository) GetTasks() ([]model.Task, error) {
	return tasks.data, nil
}

func (tasks *InMemoryTaskRepository) AddTask(task *model.Task) error {
	tasks.data = append(tasks.data, *task)
	return nil
}

func (tasks *InMemoryTaskRepository) GetTaskById(id string) (model.Task, error) {
	for _, task := range tasks.data {
		if task.ID == id {
			return task, nil
		}
	}

	return model.Task{}, errors.New("task not found")
}
