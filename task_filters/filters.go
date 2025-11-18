package taskfilters

import "github.com/Dizekee/To-do-list/tasks"

type TaskFilter interface {
	Filter(tasks []tasks.Task) []tasks.Task
}

type PriorityFilter struct {
	MinPriority int
}

func (p PriorityFilter) Filter(tasks []tasks.Task) []tasks.Task {
	var result []tasks.Task
	for _, task := range tasks {
		if task.GetPriority() >= p.MinPriority {
			result = append(result, task)
		}
	}
	return result
}
