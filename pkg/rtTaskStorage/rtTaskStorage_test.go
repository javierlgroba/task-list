package rtTaskStorage

import (
	"testing"

	"github.com/javierlgroba/task-list/pkg/task"
)

func TestAdd(t *testing.T) {
	taskStore := NewRtTaskStorage()
	task := task.Task{ID: "id", Text: "text"}
	ok := taskStore.Add(task)
	if !ok {
		t.Fail()
	}
	newTask, ok := taskStore.GetTask("id")
	if !ok {
		t.Fail()
	}
	if task != newTask {
		t.Fail()
	}
	ok = taskStore.Add(task)
	if ok {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	taskStore := NewRtTaskStorage()
	firstTask := task.Task{ID: "id", Text: "text"}
	secondTask := task.Task{ID: "id2", Text: "text"}
	taskStore.Add(firstTask)
	taskStore.Add(secondTask)
	slice := taskStore.GetAll()
	if len(slice) != 2 {
		t.Fail()
	}
	containFirst := false
	containSecond := false
	for _, containedTask := range slice {
		if containedTask == firstTask {
			containFirst = true
		}
		if containedTask == firstTask {
			containSecond = true
		}
	}
	if !containFirst || !containSecond {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	taskStore := NewRtTaskStorage()
	task := task.Task{ID: "id", Text: "text"}
	taskStore.Add(task)
	taskStore.Remove("id2")
	slice := taskStore.GetAll()
	if len(slice) != 1 {
		t.Fail()
	}
	taskStore.Remove("id")
	slice = taskStore.GetAll()
	if len(slice) != 0 {
		t.Fail()
	}
}

func TestRemoveAll(t *testing.T) {
	taskStore := NewRtTaskStorage()
	firstTask := task.Task{ID: "id", Text: "text"}
	secondTask := task.Task{ID: "id2", Text: "text"}
	taskStore.Add(firstTask)
	taskStore.Add(secondTask)
	taskStore.RemoveAll()
	slice := taskStore.GetAll()
	if len(slice) != 0 {
		t.Fail()
	}
}
