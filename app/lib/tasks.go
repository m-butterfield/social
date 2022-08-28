package lib

import (
	"bytes"
	"cloud.google.com/go/cloudtasks/apiv2"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/protobuf/types/known/durationpb"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	locationID = "us-central1"
)

type PublishPostRequest struct {
	PostID string   `json:"PostID"`
	Images []string `json:"images"`
}

type TaskCreator interface {
	CreateTask(string, string, interface{}) (*tasks.Task, error)
}

func NewTaskCreator() (TaskCreator, error) {
	workerBaseURL := os.Getenv("WORKER_BASE_URL")
	if workerBaseURL == "" {
		return nil, errors.New("WORKER_BASE_URL not set")
	}
	if strings.HasPrefix(workerBaseURL, "http://localhost") {
		return &localTaskCreator{workerBaseURL: workerBaseURL}, nil
	}
	serviceAccountEmail := os.Getenv("TASK_SERVICE_ACCOUNT_EMAIL")
	if serviceAccountEmail == "" {
		return nil, errors.New("TASK_SERVICE_ACCOUNT_EMAIL not set")
	}
	return &taskCreator{
		workerBaseURL:       workerBaseURL,
		serviceAccountEmail: serviceAccountEmail,
	}, nil
}

type localTaskCreator struct {
	workerBaseURL string
}

func (t *localTaskCreator) CreateTask(taskName, _ string, data interface{}) (*tasks.Task, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	go func() {
		_, err = http.Post(t.workerBaseURL+taskName, "application/json", bytes.NewBuffer(body))
		if err != nil {
			log.Print("Async task error:", err)
		}
	}()
	return nil, nil
}

type taskCreator struct {
	workerBaseURL       string
	serviceAccountEmail string
}

func (t *taskCreator) CreateTask(taskName, queueID string, body interface{}) (*tasks.Task, error) {
	url := t.workerBaseURL + taskName
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer func(client *cloudtasks.Client) {
		err := client.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(client)

	req := &tasks.CreateTaskRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/queues/%s", ProjectID, locationID, queueID),
		Task: &tasks.Task{
			DispatchDeadline: durationpb.New(30 * time.Minute),
			MessageType: &tasks.Task_HttpRequest{
				HttpRequest: &tasks.HttpRequest{
					HttpMethod: tasks.HttpMethod_POST,
					Url:        url,
					Headers:    map[string]string{"Content-Type": "application/json"},
					AuthorizationHeader: &tasks.HttpRequest_OidcToken{
						OidcToken: &tasks.OidcToken{
							ServiceAccountEmail: t.serviceAccountEmail,
						},
					},
				},
			},
		},
	}

	if message, err := json.Marshal(body); err != nil {
		return nil, err
	} else {
		req.Task.GetHttpRequest().Body = message
	}

	return client.CreateTask(ctx, req)
}
