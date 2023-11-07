package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reforce.pattern/config"
)

// log создание глобальной переменной для работы с логами сервиса
var log = initialize()

// initialize создание и настройка логов сервиса
func initialize() *zap.SugaredLogger {

	logWriter := zapcore.AddSync(os.Stdout)

	logCFG := zapcore.EncoderConfig{
		MessageKey:     "MESSAGE",
		LevelKey:       "LEVEL",
		TimeKey:        "TIME",
		NameKey:        "NAME",
		CallerKey:      "CALLER",
		FunctionKey:    "FUNC",
		StacktraceKey:  "STACK",
		SkipLineEnding: false,
		LineEnding:     "\n",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder

	if config.CFG.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(logCFG)
	} else {
		encoder = zapcore.NewJSONEncoder(logCFG)
	}
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(loggerLevelMap[config.CFG.Logger.Level]))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	_ = logger.Sugar().Sync()

	return logger.Sugar()
}

// For mapping config logger to email_service logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func Error(template string, args ...interface{}) {
	log.Errorf(template, args...)
}
func Debug(template string, args ...interface{}) {
	log.Debugf(template, args...)
}
func Info(template string, args ...interface{}) {
	log.Infof(template, args...)
}
func Print(template string, args ...interface{}) {
	log.Infof(template, args...)
}
func Warn(template string, args ...interface{}) {
	log.Warnf(template, args...)
}
func DPanic(template string, args ...interface{}) {
	log.DPanicf(template, args...)
}
func Panic(template string, args ...interface{}) {
	log.Panicf(template, args...)
}
func Fatal(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}
