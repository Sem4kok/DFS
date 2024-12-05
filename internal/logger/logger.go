package logger

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	With(args ...interface{}) Logger
}

type NOPLogger struct{}

func (l *NOPLogger) Debug(args ...interface{}) {}
func (l *NOPLogger) Info(args ...interface{})  {}
func (l *NOPLogger) Warn(args ...interface{})  {}
func (l *NOPLogger) Error(args ...interface{}) {}
func (l *NOPLogger) Fatal(args ...interface{}) {}
func (l *NOPLogger) With(args ...interface{}) Logger {
	return nil
}
