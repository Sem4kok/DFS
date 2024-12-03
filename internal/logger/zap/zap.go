package zap

import (
	"github.com/Sem4kok/DFS/internal/logger"
	"go.uber.org/zap"
)

// ZapLogger implements logger interface
// using sugared zap logger
type ZapLogger struct {
	Base    *zap.Logger
	Sugared *zap.SugaredLogger
}

// NewZapLoggerProd returns new ZapLogger prod-ready
func NewZapLoggerProd() *ZapLogger {
	lg, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{lg, lg.Sugar()}
}

// NewZapLoggerDev returns new ZapLogger dev-ready
func NewZapLoggerDev() *ZapLogger {
	lg, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{lg, lg.Sugar()}
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.Sugared.Debug(args)
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.Sugared.Info(args)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.Sugared.Warn(args)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.Sugared.Error(args)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.Sugared.Fatal(args)
}

func (z *ZapLogger) With(args ...interface{}) logger.Logger {
	return &ZapLogger{z.Base, z.Sugared.With(args...)}
}
