package logger

type ILogger interface {
	Log(message string)
	Error(err error)
}
