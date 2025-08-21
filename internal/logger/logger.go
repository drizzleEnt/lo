package logger

import (
	"context"
	"fmt"
)

type Logger struct {
	logChan <-chan string
}

func New(ch <-chan string) *Logger {
	return &Logger{
		logChan: ch,
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
			fmt.Println("logger get new log")
			l.WriteLog(log)
		}
	}
}

func (l *Logger) Stop() {
	fmt.Println("Logger.Stop()")
}

func (l *Logger) WriteLog(in string) {
	fmt.Println("Logger write log", in)
}
