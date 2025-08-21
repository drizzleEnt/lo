package app

import "lo/internal/logger"

func WithLogger() Option{
	return func(a *App) {
		if a.logger == nil{
			a.logger = logger.New(a.sp.LogChan())
		}
	}
}