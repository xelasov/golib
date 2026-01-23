package logger

import (
	"log"
	"os"
)

// Logger defines type for all Loggers
type Logger struct {
	loggerName string
	debugStr   string
	infoStr    string
	warnStr    string
	errStr     string
}

var (
	isDebugOn = true
	isInfoOn  = true
	isWarnOn  = true
	isErrorOn = true
)

var (
	stdOutLogger = log.New(os.Stdout, "", log.LstdFlags)
	stdErrLogger = log.New(os.Stderr, "", log.LstdFlags)
)

// NewLogger returns a new, initialized Logger instance
func NewLogger(name string) Logger {
	return Logger{loggerName: name,
		debugStr: "DEBUG: [" + name + "] ",
		infoStr:  "INFO: [" + name + "] ",
		warnStr:  "WARNING: [" + name + "] ",
		errStr:   "ERROR: [" + name + "] ",
	}
}

// IsDebugOn returns current state
func IsDebugOn() bool {
	return isDebugOn
}

// IsInfoOn returns current state
func IsInfoOn() bool {
	return isInfoOn
}

// IsWarningOn returns current state
func IsWarningOn() bool {
	return isWarnOn
}

// IsErrorOn returns current state
func IsErrorOn() bool {
	return isErrorOn
}

// SetDebug sets current state, returns previous state
func SetDebug(on bool) bool {
	prev := isDebugOn
	isDebugOn = on
	return prev
}

// SetInfo sets current state, returns previous state
func SetInfo(on bool) bool {
	prev := isInfoOn
	isInfoOn = on
	return prev
}

// SetWarning sets current state, returns previous state
func SetWarning(on bool) bool {
	prev := isWarnOn
	isWarnOn = on
	return prev
}

// SetError sets current state, returns previous state
func SetError(on bool) bool {
	prev := isErrorOn
	isErrorOn = on
	return prev
}

// Debug print info message if Info messages are on
func (l *Logger) Debug(msg string, args ...interface{}) {
	if isDebugOn {
		stdOutLogger.Printf(l.debugStr+msg, args...)
	}
}

// Info print info message if Info messages are on
func (l *Logger) Info(msg string, args ...interface{}) {
	if isInfoOn {
		stdOutLogger.Printf(l.infoStr+msg, args...)
	}
}

// Warning print info message if Info messages are on
func (l *Logger) Warning(msg string, args ...interface{}) {
	if isWarnOn {
		stdOutLogger.Printf(l.warnStr+msg, args...)
	}
}

// Error print info message if Info messages are on
func (l *Logger) Error(msg string, args ...interface{}) {
	if isErrorOn {
		stdErrLogger.Printf(l.errStr+msg, args...)
	}
}
