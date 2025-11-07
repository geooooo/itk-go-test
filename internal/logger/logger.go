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

type logger struct {
	output *os.File
}

func NewLogger(output *os.File) *logger {
	return &logger{
		output: output,
	}
}

func (l *logger) Error(err error) {
	l.log(logError, err.Error())
}

func (l *logger) Log(message string) {
	l.log(logNormal, message)
}

func (l *logger) log(level logLevel, message string) {
	prefix := "Log"
	switch level {
	case logError:
		prefix = "Error"
	}

	fmt.Fprintf(l.output, "%s: %s [%s]\n", prefix, message, time.Now().Format("2006-01-02 15:04:05"))
}
