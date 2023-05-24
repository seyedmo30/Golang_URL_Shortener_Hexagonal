package logger

import (
	"go.uber.org/zap/zapcore"
	"os"
	"sync"

	"go.uber.org/zap"
)

var once sync.Once

var instance *zap.Logger

type logger struct {
}

var logInstance logger

func Log() *logger {
	return &logInstance
}
func init() {

	once.Do(func() {

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		defaultLogLevel := zapcore.DebugLevel
		exitLevel := zapcore.FatalLevel
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)

		instance = zap.New(core, zap.AddCaller(), zap.AddStacktrace(exitLevel))

	})

}

type Interface interface {
	Debug(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Panic(message string)
	Fatal(message string)
}

var _ Interface = (*logger)(nil)

func (l *logger) Debug(message string) {
	instance.Debug(message)

}
func (l *logger) Info(message string) {
	instance.Info(message)

}

func (l *logger) Warn(message string) {
	instance.Warn(message)
}
func (l *logger) Error(message string) {
	instance.Error(message)

}
func (l *logger) Panic(message string) {
	instance.Panic(message)

}
func (l *logger) Fatal(message string) {
	instance.Fatal(message)

}
