package app

import (
	"lo/internal/controllers"
	"lo/internal/controllers/taskctrl"
)

type serviceProvider struct {
	logChan chan string

	taskctrl controllers.TaskController
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{
		logChan: make(chan string),
	}
}

func (s *serviceProvider) LogChan() chan string {
	if s.logChan == nil {
		s.logChan = make(chan string)
	}

	return s.logChan
}

func (s *serviceProvider) TaskController() controllers.TaskController {
	if s.taskctrl == nil {
		s.taskctrl = taskctrl.New()
	}

	return s.taskctrl
}
