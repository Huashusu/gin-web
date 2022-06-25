package core

import (
	"gin-web/global"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	cores := make([]zapcore.Core, 0, 7)
	d := getCore(zap.DebugLevel)
	i := getCore(zap.InfoLevel)
	w := getCore(zap.WarnLevel)
	e := getCore(zap.ErrorLevel)
	p := getCore(zap.PanicLevel)
	f := getCore(zap.FatalLevel)
	switch strings.ToLower(global.CONFIG.Log.Level) {
	case "debug":
		cores = append(cores, d, i, w, e, p, f)
	case "info":
		cores = append(cores, i, w, e, p, f)
	case "warn":
		cores = append(cores, w, e, p, f)
	case "error":
		cores = append(cores, e, p, f)
	case "panic":
		cores = append(cores, p, f)
	case "fatal":
		cores = append(cores, f)
	default:
		cores = append(cores, d, i, w, e, p, f)
	}
	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	if global.CONFIG.Log.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncodeConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     CustomEncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	if global.CONFIG.Log.Format == "json" {
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return config
}

func getEncoder() zapcore.Encoder {
	if global.CONFIG.Log.Format == "json" {
		return zapcore.NewJSONEncoder(getEncodeConfig())
	}
	return zapcore.NewConsoleEncoder(getEncodeConfig())
}

func getCore(level zapcore.Level) (core zapcore.Core) {
	writer, err := FileRotateLog.GetWriterSyncer(level.String())
	if err != nil {
		fmt.Printf("get writer syncer failed err:%v\n", err)
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

func CustomEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}
