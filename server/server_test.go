package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubTaskStore struct {
	tasks map[string]string
}

func (s *StubTaskStore) GetTask(id string) string {
	task := s.tasks[id]
	return task
}

// /tasks/all, /tasks/{id}  post: /new delete: /delete{id}
func TestGetTasks(t *testing.T) {
	store := StubTaskStore{
		map[string]string{
			"1": "walking the dog",
			"2": "coding this app",
		},
	}
	server := &TaskServer{&store}
	t.Run("returns waking the dog", func(t *testing.T) {
		request := newGetTaskRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "walking the dog"

		assertCorrectMessage(t, response.Body.String(), want)
	})
	t.Run("returns coding this app", func(t *testing.T) {
		request := newGetTaskRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := "coding this app"

		assertCorrectMessage(t, response.Body.String(), want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func newGetTaskRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/tasks/%s", id), nil)
	return req
}
