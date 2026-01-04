package logger

type LogLevel string

const (
	LogLevelWarning = "Warning"
	LogLevelError   = "Error"
	LogLevelInfo    = "Info"
)

type Logger interface {
	Log(level LogLevel, module string, msg string, args ...any)
}
