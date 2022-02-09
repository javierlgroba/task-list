package task

// task represents data about a task
type Task struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// task store is the interface to implemente by any task data storage
type TasksStore interface {
	GetAll() []Task
	GetTask(string) (Task, bool)
	Remove(string) bool
	RemoveAll() bool
	Add(Task) bool
}
