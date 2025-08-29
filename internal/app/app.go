package app

import (
	"context"
	"fmt"
	"lo/internal/domain/logs"
	"lo/internal/logger"
	"lo/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Option func(*App)

func WithLogger() Option {
	return func(a *App) {
		// if a.logger == nil{
		// 	a.logger = logger.New(a.sp.LogChan())
		// }
	}
}

type App struct {
	sp *serviceProvider

	srv *http.Server
}

func New(ctx context.Context, opts ...Option) (*App, error) {
	a := &App{}

	if err := a.initDebs(ctx); err != nil {
		return nil, err
	}

	for _, f := range opts {
		f(a)
	}

	return a, nil
}

func (a *App) TestingCalls() {
	t := time.NewTicker(1 * time.Millisecond)

	ct, c := context.WithTimeout(context.Background(), 10*time.Second)
	defer c()

l:
	for {
		select {
		case <-t.C:
			a.sp.Logger().Write(logs.LogMsg{
				Msg: time.Now().String(),
			})
		case <-ct.Done():
			break l
		}
	}

	ca, gc := logger.Stats()
	fmt.Println("Results")
	fmt.Printf("total calls: %v\n", ca)
	fmt.Printf("goroutines created: %v\n", gc)
	fmt.Println("end")
}

func (a *App) Run(ctx context.Context, cancel context.CancelFunc) error {
	// a.TestingCalls()
	// return nil

	go func() {
		if err := a.runHttpServer(); err != nil {
			if err == http.ErrServerClosed {
				return
			}
			log.Printf("failed to run HTTP server: %v", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal)

	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down http server...")
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Printf("failed Shutting down http server %s.\n", err.Error())
		return err
	}
	log.Println("http server stopped")

	return nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initHttpSrv,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.sp = newServiceProvider()
	return nil
}

func (a *App) initHttpSrv(ctx context.Context) error {
	srv := &http.Server{
		Addr:           "localhost:8080",
		Handler:        routes.InitRoutes(a.sp.TaskController(), a.sp.Logger()),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	a.srv = srv

	return nil
}

func (a *App) runHttpServer() error {
	log.Printf("server run on :8080")
	err := a.srv.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
