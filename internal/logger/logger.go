package logger

type Logger interface {
	Log(message string)
	Error(err error)
}
