package manager

import (
	"fmt"
	"time"

	"github.com/Dizekee/To-do-list/tasks"
)

type TaskManager struct {
	tasks  []tasks.Task
	nextID int
}

func (tm *TaskManager) AddTask(taskType, title, description string, priority int, dueDate time.Time, extraParams map[string]string) error {
	basicTask := tasks.BasicTask{
		Id:          tm.nextID,
		Title:       title,
		Description: description,
		Complited:   false,
		Priority:    priority,
		DueDate:     dueDate,
	}

	if priority < 1 || priority > 5 {
		return fmt.Errorf("приоритет должен быть от 1 до 5")
	}

	var task tasks.Task

	switch taskType {
	case "work":
		task = &tasks.WorkTask{
			BasicTask: basicTask,
			Project:   extraParams["project"],
			Assignee:  extraParams["extraparams"],
		}
	case "personal":
		task = &tasks.PersonalTask{
			BasicTask: basicTask,
			Category:  extraParams["category"],
			Location:  extraParams["location"],
		}
	}
	tm.tasks = append(tm.tasks, task)
	tm.nextID++
	return nil
}

func (tm *TaskManager) RemoveTask(id int) bool {
	for _, task := range tm.tasks {
		if task.GetID() == id {
			tm.tasks = append(tm.tasks[:id], tm.tasks[id+1:]...)
			return true
		}
	}
	return false
}

func (tm *TaskManager) GetTask(id int) (tasks.Task, error) {
	for _, task := range tm.tasks {
		if task.GetID() == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("Задача по ID %d не найдена", id)
}

func (tm *TaskManager) CompleteTask(id int) bool {
	for _, task := range tm.tasks {
		if task.GetID() == id {
			task.MarkComplited()
			return true
		}
	}
	return false
}

func (tm *TaskManager) GetOverdueTasks() []tasks.Task {
	var overrideTasks []tasks.Task
	for _, task := range tm.tasks {
		if !task.IsComplited() && task.DaysUntilDue() < 0 {
			overrideTasks = append(overrideTasks, task)
		}
	}
	return overrideTasks
}

func (tm *TaskManager) GetTasksByPriority(minPriority int) []tasks.Task {
	var priorityTasks []tasks.Task
	for _, task := range tm.tasks {
		if task.GetPriority() >= minPriority && !task.IsComplited() {
			priorityTasks = append(priorityTasks, task)
		}
	}
	return priorityTasks
}

func (tm *TaskManager) GetTasksByType(taskType string) []tasks.Task {
	var filteredTasks []tasks.Task
	for _, task := range tm.tasks {
		switch taskType {
		case "work":
			filteredTasks = append(filteredTasks, task)
		case "personal":
			filteredTasks = append(filteredTasks, task)
		}
	}
	return filteredTasks
}
