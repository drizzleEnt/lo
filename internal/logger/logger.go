package logger

import (
	"context"
	"fmt"
	"lo/internal/domain/logs"
	"sync"
)

type Option func(*Logger)

func WithWorkers(workers int) Option {
	return func(l *Logger) {
		l.workers = workers
	}
}

func WithBuffer(buffer int) Option {
	return func(l *Logger) {
		l.buffer = buffer
	}
}

type Logger struct {
	logChan  chan logs.LogMsg
	stopChan chan struct{}
	workers  int
	buffer   int
	wg       *sync.WaitGroup
}

func New(opts ...Option) *Logger {
	return &Logger{
		logChan:  make(chan logs.LogMsg),
		stopChan: make(chan struct{}),
		wg:       &sync.WaitGroup{},
	}
}

func (l *Logger) Run(ctx context.Context) error {
	fmt.Println("logger run")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("logger ctx canceled")
			l.Stop()
			return nil
		case log := <-l.logChan:
			fmt.Println("logger get new log", log)
		}
	}
}

func (l *Logger) Stop() {
	fmt.Println("Logger.Stop()")
}

func (l *Logger) Write(in logs.LogMsg) {
	go func() {
		l.logChan <- in
	}()
}
