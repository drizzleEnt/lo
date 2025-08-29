package routes

import (
	"lo/internal/controllers"
	"lo/internal/middlerware"
	"net/http"
)

func InitRoutes(c controllers.TaskController, logger middlerware.Logger) http.Handler {
	mux := http.NewServeMux()

	logMW := middlerware.Logging(logger)
	//errMW := middlerware.ErrorLogging(logger)

	common := withCommon(logMW)

	mux.Handle("/tasks", adapt(c.GetTasks, common...))

	mux.Handle("/task/", adapt(func(w http.ResponseWriter, r *http.Request) error {
		switch r.Method {
		case http.MethodGet:
			return c.GetTaskById(w, r)
		case http.MethodPost:
			return c.CreateTask(w, r)
		default:
			ErrorMethod(w, r)
		}
		return nil
	}, common...))

	return mux
}

func ErrorMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

type Middleware func(next middlerware.HandlerWithErr) middlerware.HandlerWithErr

func adapt(h middlerware.HandlerWithErr, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func withCommon(mw ...Middleware) []Middleware {
	return mw
}
