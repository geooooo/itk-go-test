package logger

import (
	"fmt"
	"os"
	"time"
)

type logLevel int

const (
	logNormal = logLevel(iota)
	logError
)

type Logger struct {
	output *os.File
}

func NewLogger(output *os.File) *Logger {
	return &Logger{
		output: output,
	}
}

func (l *Logger) Error(err error) {
	l.log(logError, err.Error())
}

func (l *Logger) Log(message string) {
	l.log(logNormal, message)
}

func (l *Logger) log(level logLevel, message string) {
	prefix := "Log"
	switch level {
	case logError:
		prefix = "Error"
	}

	fmt.Fprintf(l.output, "%s: %s [%s]\n", prefix, message, time.Now())
}
