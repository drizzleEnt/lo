package logs

type LogMsg struct {
	Level LogLevel
	Msg   string
	Err   error
}

type LogLevel string

const (
	InfoLogLevel  LogLevel = "[INFO]"
	ErrorLogLevel LogLevel = "[ERROR]"
	DebugLogLevel LogLevel = "[DEBUG]"
)
