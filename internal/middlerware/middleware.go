package middlerware

import (
	"lo/internal/domain/logs"
	"net/http"
)

type Logger interface {
	Write(logs.LogMsg)
}

type HandlerWithErr func(w http.ResponseWriter, r *http.Request) error

func InfoLog(logger Logger, r *http.Request) func(r *http.Request) {
	return func(r *http.Request) {
		logger.Write(logs.LogMsg{
			Level: logs.InfoLogLevel,
			Msg:   "Income request" + r.Method + r.URL.Path,
		})
	}
}

func Logging(logger Logger) func(next HandlerWithErr) HandlerWithErr {
	return func(next HandlerWithErr) HandlerWithErr {
		return func(w http.ResponseWriter, r *http.Request) error {
			//now := time.Now()

			if err := next(w, r); err != nil {

			}

			logger.Write(logs.LogMsg{
				Level: logs.InfoLogLevel,
				Msg:   "Income request " + r.Method + " " + r.URL.Path,
			})
			return next(w, r)
		}
	}
}

func ErrorLogging(logger Logger) func(next HandlerWithErr) HandlerWithErr {
	return func(next HandlerWithErr) HandlerWithErr {
		return func(w http.ResponseWriter, r *http.Request) error {
			err := next(w, r)
			if err != nil {
				logger.Write(logs.LogMsg{
					Level: logs.ErrorLogLevel,
					Msg:   "error in request" + r.Method + " " + r.URL.Path,
					Err:   err,
				})
			}
			return err
		}
	}
}
