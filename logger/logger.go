package logger

import (
	"fmt"
	"sync"
)

var (
	defaultLogger Logger
	mu            sync.RWMutex
)

func SetLogger(l Logger) {
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = l
}

func Log(level LogLevel, module string, msg string, args ...any) {
	mu.RLock()
	defer mu.RUnlock()
	if defaultLogger != nil {
		defaultLogger.Log(level, module, msg, args...)
	}
}

// DEFAULT LOGGER STDOUT
type StdLogger struct{}

func (l *StdLogger) Log(level LogLevel, module string, msg string, args ...any) {
	fmt.Printf("[%s] [%s] %s\n", level, module, fmt.Sprintf(msg, args...))
}
