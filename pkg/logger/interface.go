package logger

type ILogger interface {
	Info(msg string, meta interface{})
	Warn(msg string, meta interface{})
	Debug(msg string, meta interface{})
	MError(msg string, meta interface{})
}
