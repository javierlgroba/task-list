package rtTaskStorage

import (
	"sort"

	"github.com/javierlgroba/task-list/pkg/task"
)

type RtTaskStorage struct {
	tasks map[string]task.Task
}

func NewRtTaskStorage() RtTaskStorage {
	var m RtTaskStorage
	m.tasks = make(map[string]task.Task)
	return m
}

func (r RtTaskStorage) GetAll() []task.Task {
	keys := make([]string, 0, len(r.tasks))
	for k := range r.tasks {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	v := make([]task.Task, 0, len(r.tasks))
	for _, k := range keys {
		v = append(v, r.tasks[k])
	}
	return v
}

func (r RtTaskStorage) GetTask(s string) (task.Task, bool) {
	val, ok := r.tasks[s]
	return val, ok
}

func (r RtTaskStorage) Remove(s string) bool {
	val, ok := r.tasks[s]
	if ok {
		delete(r.tasks, val.ID)
	}
	return ok
}

func (r RtTaskStorage) RemoveAll() bool {
	for k := range r.tasks {
		delete(r.tasks, k)
	}
	return true
}

func (r RtTaskStorage) Add(t task.Task) bool {
	_, ok := r.tasks[t.ID]
	if !ok {
		r.tasks[t.ID] = t
	}
	return !ok
}
