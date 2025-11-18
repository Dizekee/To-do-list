package tasks

type WorkTask struct {
	BasicTask
	Project  string
	Assignee string
}

// Добавьте методы специфичные для WorkTask если нужно
func (w *WorkTask) GetProject() string  { return w.Project }
func (w *WorkTask) GetAssignee() string { return w.Assignee }
