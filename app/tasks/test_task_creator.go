package tasks

import "google.golang.org/genproto/googleapis/cloud/tasks/v2"

type TestTaskCreator struct {
	TestCreateTask      func(string, string, interface{}) (*tasks.Task, error)
	CreateTaskCallCount int
}

func (t *TestTaskCreator) CreateTask(taskName, queueID string, body interface{}) (*tasks.Task, error) {
	t.CreateTaskCallCount += 1
	return t.TestCreateTask(taskName, queueID, body)
}
