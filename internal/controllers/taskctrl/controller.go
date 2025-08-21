package taskctrl

import "net/http"

func New() *controller {
	return &controller{}
}

type controller struct {
}

// CreateTask implements controllers.TaskController.
func (c *controller) CreateTask(w http.ResponseWriter, r *http.Request) error {
	panic("unimplemented")
}

// GetTaskById implements controllers.TaskController.
func (c *controller) GetTaskById(w http.ResponseWriter, r *http.Request) error {
	panic("unimplemented")
}

// GetTasks implements controllers.TaskController.
func (c *controller) GetTasks(w http.ResponseWriter, r *http.Request) error {
	panic("unimplemented")
}
