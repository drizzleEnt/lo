package app

import (
	"lo/internal/controllers"
	"lo/internal/controllers/taskctrl"
	"lo/internal/logger"
)

type serviceProvider struct {
	logger *logger.Logger

	taskctrl controllers.TaskController
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) Logger() *logger.Logger {
	if s.logger == nil {
		s.logger = logger.New(
			logger.WithSemaphore(100),
		)
	}

	return s.logger
}

func (s *serviceProvider) TaskController() controllers.TaskController {
	if s.taskctrl == nil {
		s.taskctrl = taskctrl.New()
	}

	return s.taskctrl
}
