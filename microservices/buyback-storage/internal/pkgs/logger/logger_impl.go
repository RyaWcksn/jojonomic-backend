package logger

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type runtimeCaller struct {
	function string
	fileName string
	line     int
}

type Logger struct {
	log     *zap.Logger
	service string
	env     string
}

var Log *Logger

// Initialize Logger instance.
func New(service, env, logLevel string) *Logger {
	Log = Init(service, env, logLevel)
	return Log
}

func (l *Logger) Debug(msg string) {
	debugInfo := getLogInfo()
	l.log.Debug(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", debugInfo.fileName),
		zap.String("function", debugInfo.function),
		zap.Int("line", debugInfo.line),
	)
}

func (l *Logger) Info(msg string) {
	runtime := getLogInfo()
	l.log.Info(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Warn(msg string) {
	runtime := getLogInfo()
	l.log.Warn(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Error(msg string) {
	runtime := getLogInfo()
	l.log.Error(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Fatal(msg string) {
	runtime := getLogInfo()
	l.log.Fatal(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Panic(msg string) {
	runtime := getLogInfo()
	l.log.Panic(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Debugf(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	runtime := getLogInfo()

	l.log.Debug(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Infof(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	runtime := getLogInfo()

	l.log.Info(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Warnf(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	runtime := getLogInfo()

	l.log.Warn(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Errorf(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	runtime := getLogInfo()

	l.log.Error(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	runtime := getLogInfo()

	l.log.Fatal(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

func (l *Logger) Panicf(format string, args ...interface{}) {

	msg := fmt.Sprintf(format, args...)

	runtime := getLogInfo()

	l.log.Panic(
		msg,
		zap.String("service", l.service),
		zap.String("env", l.env),
		zap.String("file", runtime.fileName),
		zap.String("function", runtime.function),
		zap.Int("line", runtime.line),
	)
}

// getLogInfo return function, file, and line for logging purpose
func getLogInfo() runtimeCaller {
	pc, file, line, _ := runtime.Caller(2)
	return runtimeCaller{
		function: runtime.FuncForPC(pc).Name(),
		fileName: file,
		line:     line,
	}
}

func Init(service, env, logLevel string) *Logger {

	if logLevel == "" {
		logLevel = "error"
	}

	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		return nil
	}

	logger, err := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}.Build()
	if err != nil {
		return nil
	}
	defer logger.Sync()

	return &Logger{
		log:     logger,
		service: service,
		env:     env,
	}
}
