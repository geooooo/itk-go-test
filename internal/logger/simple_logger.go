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

type simpleLogger struct {
	output *os.File
}

func NewSimpleLogger(output *os.File) *simpleLogger {
	return &simpleLogger{
		output: output,
	}
}

func (l *simpleLogger) Error(err error) {
	l.log(logError, err.Error())
}

func (l *simpleLogger) Log(message string) {
	l.log(logNormal, message)
}

func (l *simpleLogger) log(level logLevel, message string) {
	prefix := "Log"
	switch level {
	case logError:
		prefix = "Error"
	}

	fmt.Fprintf(l.output, "%s: %s [%s]\n", prefix, message, time.Now())
}
