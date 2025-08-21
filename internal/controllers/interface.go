package controllers

import "net/http"

type TaskController interface {
	GetTasks(w http.ResponseWriter, r *http.Request) error
	GetTaskById(w http.ResponseWriter, r *http.Request) error
	CreateTask(w http.ResponseWriter, r *http.Request) error
}
