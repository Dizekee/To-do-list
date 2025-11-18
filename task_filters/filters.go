package taskfilters

import "github.com/Dizekee/To-do-list/tasks"

type TaskFilter interface {
	Filter(tasks []tasks.Task) []tasks.Task
}

type PriorityFilter struct {
	MinPriority int
}
