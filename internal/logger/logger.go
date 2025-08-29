package logger

import (
	"context"
	"fmt"
	"lo/internal/domain/logs"
	"sync"
	"time"
)

var m *sync.Mutex = &sync.Mutex{}
var calls int

var mu *sync.Mutex = &sync.Mutex{}
var goroutinesCalls int

type Option func(*Logger)

// With losing data but without freeze
func WithSemaphore(buffer int) Option {
	return func(l *Logger) {
		l.buffer = buffer
		l.sem = make(chan struct{}, buffer)
	}
}

type Logger struct {
	sem    chan struct{}
	buffer int
}

func New(opts ...Option) *Logger {
	l := &Logger{}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *Logger) Run(ctx context.Context) error {
	fmt.Println("logger run")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("logger ctx canceled")
			l.Stop()
			return nil
		default:
			continue
		}
	}
}

func (l *Logger) Stop() {
	fmt.Println("Logger.Stop()")
}

func (l *Logger) Write(in logs.LogMsg) {
	l.withSem(func() {
		l.saveLog(in)
	})
}

func (l *Logger) withSem(f func()) {
	if l.sem == nil {
		f()
		return
	}

	m.Lock()
	calls++
	m.Unlock()

	select {
	case l.sem <- struct{}{}:
	default:
		return
	}

	go func() {
		f()

		<-l.sem
	}()
	mu.Lock()
	goroutinesCalls++
	mu.Unlock()
}

func (l *Logger) saveLog(log logs.LogMsg) {
	fmt.Println("logger get new log", log)
	time.Sleep(1 * time.Second)
	fmt.Println("logger save log in storage", log)
}

func Stats() (int, int) {
	return calls, goroutinesCalls
}
