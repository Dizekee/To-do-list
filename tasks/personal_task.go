package tasks

type PersonalTask struct {
	BasicTask
	Category string
	Location string
}

// Добавьте методы специфичные для PersonalTask если нужно
func (p *PersonalTask) GetCategory() string { return p.Category }
func (p *PersonalTask) GetLocation() string { return p.Location }
