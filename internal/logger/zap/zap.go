package zap

import (
	"go.uber.org/zap"
)

// ZapLogger implements logger interface
// using sugared zap logger
type ZapLogger struct {
	*zap.SugaredLogger
}

// NewZapLogger returns new ZapLogger
func NewZapLogger() *ZapLogger {
	lg, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{lg.Sugar()}
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.Debug(args)
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.Info(args)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.Warn(args)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.Info(args)
}
