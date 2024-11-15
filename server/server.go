package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Task struct {
	id       int
	title    string
	time     string
	date     string
	taskType string
}

type TaskStore interface {
	GetTask(id string) string
}

type TaskServer struct {
	store TaskStore
}

func (t *TaskServer) ServeHttp(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	fmt.Fprintf(w, t.store.GetTask(id))
}
