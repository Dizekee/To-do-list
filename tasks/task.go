package tasks

import "time"

type Task interface {
	GetID() int
	GetTitle() string
	GetDescription() string
	IsComplited() bool
	MarkComplited()
	GetPriority() int
	GetDueDate() time.Time
	DaysUntilDue() int
}

type BasicTask struct {
	Id          int
	Title       string
	Description string
	Complited   bool
	Priority    int
	DueDate     time.Time
}

func (b *BasicTask) GetID() int             { return b.Id }
func (b *BasicTask) GetTitle() string       { return b.Title }
func (b *BasicTask) GetDescription() string { return b.Description }
func (b *BasicTask) IsComplited() bool      { return b.Complited }
func (b *BasicTask) MarkComplited()         { b.Complited = true }
func (b *BasicTask) GetPriority() int       { return b.Priority }
func (b *BasicTask) GetDueDate() time.Time  { return b.DueDate }
func (b *BasicTask) DaysUntilDue() int {
	return int(time.Until(b.DueDate).Hours() / 24)
}
