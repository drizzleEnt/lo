package taskctrl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func New() *controller {
	return &controller{}
}

type controller struct {
}

// CreateTask implements controllers.TaskController.
func (c *controller) CreateTask(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("CreateTask")
	b := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&b)
	if _, ok := b["err"]; ok {
		return fmt.Errorf("CreateTask some error")
	}
	return nil
}

// GetTaskById implements controllers.TaskController.
func (c *controller) GetTaskById(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("GetTaskByID")
	ok := r.URL.Query().Get("err")
	if ok != "" {
		return fmt.Errorf("GetTaskById some error")
	}

	return nil
}

// GetTasks implements controllers.TaskController.
func (c *controller) GetTasks(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("TaskController.GetTasks")
	ok := r.URL.Query().Get("err")
	if ok != "" {
		return fmt.Errorf("GetTasks some error")
	}

	time.Sleep(5 * time.Second)

	return nil
}
