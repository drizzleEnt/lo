package routes

import (
	"lo/internal/controllers"
	"net/http"
)

func InitRoutes(c controllers.TaskController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if err := c.GetTasks(w, r); err != nil {

			}
		default:
			ErrorMethod(w, r)
		}
	})

	mux.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if err := c.GetTaskById(w, r); err != nil {

			}
		case http.MethodPost:
			if err := c.CreateTask(w, r); err != nil {

			}
		default:
			ErrorMethod(w, r)
		}
	})

	return mux
}

func AuthMiddleware() {

}

func LogMiddleware() {

}

func ErrorMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
